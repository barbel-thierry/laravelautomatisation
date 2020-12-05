# laravelautomatisation

## Install

```
go get -u github.com/barbel-thierry/laravelautomatisation
```

For comfort purpose, add these aliases in ~/.bashrc:

```
alias lpn="go run github.com/barbel-thierry/laravelautomatisation new"
alias lpo="go run github.com/barbel-thierry/laravelautomatisation open"
alias lpr="go run github.com/barbel-thierry/laravelautomatisation reset docker_db_1"
```

## Usage

* `new` takes as extra parameter the name of the project you want to initiate and the version of Laravel you want to use (if defers from the current)
* `reset` takes as extra parameter the name of your mysql docker container ('docker_db_1' is the default name)
