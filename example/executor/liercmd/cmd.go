package main

import (
	"fmt"
	"github.com/aide-cloud/universal/executor"
)

type MyServer struct{}
type ChildServer struct{}

func (m *ChildServer) Start() error {
	// do something
	fmt.Println("ChildServer start")
	return nil
}

func (m *ChildServer) Stop() {
	// do something
	fmt.Println("ChildServer stop")
}

func NewChildServer() *ChildServer {
	return &ChildServer{}
}

func (m *MyServer) Start() error {
	// do something
	fmt.Println("my server start")
	return nil
}

func (m *MyServer) Stop() {
	// do something
	fmt.Println("my server stop")
}

func NewMyServer() *MyServer {
	return &MyServer{}
}

func main() {
	executor.ExecMulSerProgram(executor.NewLierCmd(
		executor.WithServiceName("master"),
		executor.WithProperty(map[string]string{
			"version": "1.0.0",
			"name   ": "master",
			"time   ": "2020-12-12 12:12:12",
			"author ": "aide-cloud",
		}),
		executor.WithServices(NewMyServer(), NewChildServer()),
	))
}
