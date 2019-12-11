# Subprocess

A subprocess API similar to that of `python` `subprocess` module



### Installing

Get the module using

	```
	go get github.com/millefalcon/Go-Subprocess/subprocess
	```

### Usage


```
package main

import (
        "fmt"
)

import "github.com/millefalcon/Go-Subprocess/subprocess"

func main() {
        proc := subprocess.Popen("ls")
        fmt.Println(proc.Stdout.Read())
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


