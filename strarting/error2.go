package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	uid := "yrueg"
	pwd := "353js"
	file := "mapangale"
	loginAndGetData(uid, pwd, file)
	err := fileChecker("notthere.tx")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File does not exist")
		}
	}
}

type Status int

const (
	Invalidlogin = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	message string
}

func getData(file string) ([]byte, error) {
	return nil, nil
}

func (se StatusErr) Error() string {
	return se.message
}

func login(uid, pwd string) error {
	return nil
}

func loginAndGetData(uid, pwd, file string) ([]byte, error) {
	err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  Invalidlogin,
			message: fmt.Sprintf("Invalid credentials for user %d", uid),
		}

	}
	data, err := getData(file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			message: fmt.Sprintf("File not found %s", file),
		}
	}
	return data, nil
}

//wrapping errors-When you preserve an error while adding
//additional information, it is called wrapping the erro

// is and as errors
// errrs.is() - takes two parameters error being checked and intance you are comparing it with

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		fmt.Println("Not here: %w", err)
	}
	f.Close()
	return nil
}
