package subproc

import (
//	"fmt"
	"bytes"
	"os/exec"
)

type Std_out_err_t struct {
	data bytes.Buffer
}

func (s Std_out_err_t) Read() string {
	return string(s.data.Bytes())
}

type Process struct {
	Proc *exec.Cmd
	Err error
	Stdout Std_out_err_t
	Stderr Std_out_err_t
}

func Popen(cmd ...string) Process {
	proc := exec.Command(cmd[0], cmd[1:]...)
	var stdout, stderr bytes.Buffer
	proc.Stdout = &stdout
	proc.Stderr = &stderr

	err := proc.Run()
	/*
	if err != nil {
		//fmt.Printf("cmd.Run() failed with %s\n", err)
		panic(err)
	}*/
	return Process{proc, err, Std_out_err_t{stdout}, Std_out_err_t{stderr}}
	//outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	//fmt.Println("out:", outStr, "err:", errStr)

}
