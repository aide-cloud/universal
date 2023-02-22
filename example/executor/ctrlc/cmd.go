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
	fmt.Println("start")
	return nil
}

func (m *MyServer) Stop() {
	// do something
	fmt.Println("stop")
}

func (m *MyServer) ServicesRegistration() []executor.Service {
	return []executor.Service{
		NewChildServer(),
	}
}

func NewMyServer() *MyServer {
	return &MyServer{}
}

func main() {
	executor.ExecMulSerProgram(NewMyServer())
}
