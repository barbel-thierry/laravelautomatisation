# laravelautomatisation

## Install

```
go get github.com/barbel-thierry/laravelautomatisation
```

For comfort purpose, add these aliases in ~/.bashrc:

```
alias lpn="go run github.com/barbel-thierry/laravelautomatisation new"
alias lpo="go run github.com/barbel-thierry/laravelautomatisation open"
alias lpr="go run github.com/barbel-thierry/laravelautomatisation reset docker_db_1"
```

## Usage

* `new` takes as extra parameter the name of the project you want to initiate
* `reset` takes as extra parameter the name of your mysql docker container ('docker_db_1' is the default name)