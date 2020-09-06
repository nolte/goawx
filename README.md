# goawx

AWX (Ansible Tower) SDK for the Go programming language.

Please consider this code as beta right now. It could break, but I'll try my best to keep breaking changes out of the
`master` branch. I'm working on my system tests for all endpoints.

This SDK has been developed against AWX `14.0.0`.

## Installing

```
go get -u github.com/mrcrilly/goawx
```

## Example

We can simply import `goawx` and call its services, such as the PingService:

```
import (
    "log"
    "github.com/mrcrilly/goawx/client"
)

func main() {
    client := awx.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := client.PingService.Ping()
    if err != nil {
        log.Fatalf("Ping awx err: %s", err)
    }

    log.Println("Ping awx: ", result)
}
```

More examples can be found at [here](https://github.com/mrcrilly/goawx/tree/master/examples).

## Roadmap

goawx is still in development, and its roadmap could be found at [here](https://github.com/mrcrilly/goawx/blob/master/ROADMAP.md).

## Contribute

There are many ways to contribute to awx-go.

* Submit bugs via [Github issues](https://github.com/mrcrilly/goawx/issues);
* Submit a [pull request](https://github.com/mrcrilly/goawx/pulls) for fixes or features;

### Local Run

Using [nektos/act](https://github.com/nektos/act) for start the GitHub Workflow locally.

```sh
act -j build -P ubuntu-latest=nektos/act-environments-ubuntu:18.0
```
