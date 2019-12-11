package subprocess

import (
//	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
)

type Stdout_t struct {
	stdout io.ReadCloser
}

type Process struct {
	Proc *exec.Cmd
	Stdout Stdout_t
}

func (p Process) Wait() {
	p.Wait()
}

func (s Stdout_t) Read() string {
	data, _ := ioutil.ReadAll(s.stdout)
	return string(data)
}


func Popen(cmd ...string) Process {
	proc := exec.Command(cmd[0], cmd[1:]...)
	//fmt.Printf("%T\n", proc)
	stdout, _ := proc.StdoutPipe()
	//fmt.Printf("%T\n", stdout)
	proc.Start()
	//return proc, Stdout{stdout}
	return Process{proc, Stdout_t{stdout}}
}
