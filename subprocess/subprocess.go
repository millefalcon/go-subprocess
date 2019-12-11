package subprocess

import (
//	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
)

type Std_out_err_t struct {
	data io.ReadCloser
}

type Process struct {
	Proc *exec.Cmd
	Stdout Std_out_err_t
	Stderr Std_out_err_t
}

func (p Process) Wait() {
	p.Wait()
}

func (s Std_out_err_t) Read() string {
	data, _ := ioutil.ReadAll(s.data)
	return string(data)
}


func Popen(cmd ...string) Process {
	proc := exec.Command(cmd[0], cmd[1:]...)
	stdout, _ := proc.StdoutPipe()
	stderr, _ := proc.StderrPipe()
	proc.Start()
	return Process{proc, Std_out_err_t{stdout}, Std_out_err_t{stderr}}
}
