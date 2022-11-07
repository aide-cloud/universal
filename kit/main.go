package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

var workerMode = flag.String("w", "help", "worker mode")
var repo = flag.String("r", "https://github.com/aide-cloud/aide-family-layout.git", "layout repo")
var repoPath = flag.String("p", "", "repo path")
var nomod = flag.Bool("n", false, "no mod")

// 是否强制覆盖
var force = flag.Bool("f", false, "force to overwrite")

var repoAddIgnores = []string{
	"README.md", "LICENSE", ".gitignore", "cmd", "configs", "internal", "Makefile", "Dockerfile",
}

var moduleAddIgnores = []string{
	"go.mod", "go.sum",
}

func main() {
	flag.Parse()
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

// gitClone clones a git repository to the given directory.
func gitClone(dir, repo string) {
	// bash -c "git clone repo --depth 1 dir"
	_, err := runCommand("", "bash", "-c", fmt.Sprintf("git clone %s %s", repo, path.Join("/tmp", dir)))
	if err != nil {
		log.Println("git clone failed", err)
	}

	_ = os.Mkdir(path.Join(dir), os.ModePerm)

	// copy dir to current dir
	// cp repoAddIgnores dir
	for _, ignore := range repoAddIgnores {
		_, err = runCommand("", "cp", "-r", path.Join("/tmp", dir, ignore), path.Join("./", dir, ignore))
		if err != nil {
			log.Println("cp failed", err)
		}
	}

	// nomod
	if !*nomod {
		for _, ignore := range moduleAddIgnores {
			_, err = runCommand("", "rm", "-rf", path.Join("./", dir, ignore))
			if err != nil {
				log.Println("rm failed", err)
			}
		}
	}

	// rm -rf dir
	_, err = runCommand("", "rm", "-rf", path.Join("/tmp", dir))
	if err != nil {
		log.Println("rm failed", err)
	}
}
