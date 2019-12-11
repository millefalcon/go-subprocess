package main

import (
        "fmt"
)

import "github.com/millefalcon/subprocess/go-subprocess/subproc"

func main() {
	proc := subproc.Popen("ls", "-l")
	if err := proc.Err; err != nil {
		panic(err)
	}
	fmt.Println("stdout:", proc.Stdout.Read())
	//fmt.Println("stderr:", proc.Stderr.Read())
}
