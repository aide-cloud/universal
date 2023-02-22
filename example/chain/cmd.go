package main

import (
	"errors"
	"fmt"
	"github.com/aide-cloud/universal/chain"
)

func Action1() error {
	// do something
	return nil
}

func Action2() error {
	// do something
	return nil
}

func Action3() error {
	// do something
	return nil
}

func CheckName(name string) error {
	if name == "" {
		return errors.New("name is empty")
	}

	if len(name) > 10 {
		return errors.New("name is too long")
	}

	return nil
}

func CheckAge(age int) error {
	if age < 0 {
		return errors.New("age is too small")
	}

	if age > 100 {
		return errors.New("age is too big")
	}

	return nil
}

func main() {
	chainAction := chain.NewChain(
		chain.WithTask(
			Action1,
			Action2,
			Action3,
			func() error {
				return CheckName("aide")
			},
			func() error {
				return CheckAge(20)
			},
		),
	)
	if err := chainAction.Do(); err != nil {
		panic(err)
	}

	fmt.Println("done")
}
