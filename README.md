# Subprocess

A subprocess API similar to that of `python` `subprocess` module



### Installing

Get the module using

	`go get github.com/millefalcon/go-subprocess/subprocess`

or the other one

	`go get github.com/millefalcon/go-subprocess/subproc`

### Usage


```golang
package main

import (
        "fmt"
)

import "github.com/millefalcon/go-subprocess/subprocess"

func main() {
        proc := subprocess.Popen("ls")
        fmt.Println(proc.Stdout.Read())
}
```

Alternate implementation:

```golang
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
```


## Running the tests

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
```

## Deployment

Add additional notes about how to deploy this on a live system


## Authors

* **Hanish K H** - *Initial work* - [millefalcon](https://github.com/millefalcon)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration
* etc


