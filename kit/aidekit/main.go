package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

var workerMode = flag.String("w", "help", "worker mode")
var repo = flag.String("r", "https://github.com/aide-cloud/aide-family-layout.git", "layout repo")
var repoPath = flag.String("p", "", "repo path")
var nomod = flag.Bool("n", false, "no mod")
var version = flag.Bool("v", false, "version")
var v = "v1.1.3"

var moduleAddIgnores = []string{
	"go.mod", "go.sum",
}

func main() {
	flag.Parse()
	if *version {
		fmt.Println(v)
		return
	}
	switch *workerMode {
	case "help":
		flag.Usage()
	case "new":
		gitClone(*repoPath, *repo)
	}
}

func runCommand(path, name string, arg ...string) (msg string, err error) {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	cmd.Dir = path
	err = cmd.Run()
	log.Println(cmd.Args)
	if err != nil {
		msg = fmt.Sprint(err) + ": " + stderr.String()
		err = errors.New(msg)
		log.Println("err", err.Error(), "cmd", cmd.Args)
	}
	log.Println(out.String())
	return
}

// getModuleName 获取本地module名称（go.mod第一行）
func getModuleName() string {
	cmd := exec.Command("go", "list", "-m")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		// 获取运行的根目录
		dir, _ := os.Getwd()
		return path.Base(dir)
	}
	str := out.String()
	return str[:len(str)-1]
}

// gitClone clones a git repository to the given directory.
func gitClone(filePath, repo string) {
	dir := path.Join("./", filePath)
	tmpPath := "tmp"

	var err error
	// 获取本地module名称（go.mod第一行），用于替换
	moduleName := path.Join(getModuleName(), dir)

	_ = os.RemoveAll(dir)

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return
	}

	err = os.RemoveAll(tmpPath)
	if err != nil {
		return
	}

	err = os.Mkdir(tmpPath, os.ModePerm)
	if err != nil {
		_ = os.RemoveAll(path.Base(dir))
		return
	}

	// bash -c "git clone repo --depth 1 dir"
	_, err = runCommand("", "bash", "-c", fmt.Sprintf("git clone %s --depth=1 %s", repo, path.Join(tmpPath, path.Base(dir))))
	if err != nil {
		log.Println("创建项目失败")
		_ = os.RemoveAll(path.Base(dir))
		_ = os.RemoveAll(tmpPath)
		return
	}

	_ = os.RemoveAll(path.Join(tmpPath, path.Base(dir), ".git"))

	_, _ = runCommand("", "cp", "-R", path.Join(tmpPath, path.Base(dir)), path.Dir(dir))

	_ = os.RemoveAll(tmpPath)

	// nomod
	if !*nomod {
		for _, ignore := range moduleAddIgnores {
			_ = os.Remove(path.Join(dir, ignore))
		}
	}

	moduleName = strings.ReplaceAll(moduleName, "/", "\\/")

	// 遍历dir下所有文件，替换module名称
	_ = filepath.Walk(path.Join("./", dir), func(path string, info fs.FileInfo, err error) error {
		if info == nil || info.IsDir() {
			return nil
		}
		_, _ = runCommand("", "sed", "-i", "", fmt.Sprintf("s/github.com\\/aide-cloud\\/aide-family-layout/%s/g", moduleName), path)
		return nil
	})
}
