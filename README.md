## Install
As this is a private repository on Github you need to create the project's folder in your GOPATH src folder.
Create a folder under:
```
[GOPATH]/src/github.com/capsulemaglev/league_backend_challenge
```
Access the folder and proceed to either build or run.

## How to build

1. Simply run:

```
go build ./cmd/matrix-server/main.go
```

And a executable will be built in the working directory.

## How to run

1. Run web server:

```
go run ./cmd/matrix-server/main.go
```

2. Send request:

```
curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
```

Valid endpoints are:

```
/echo
/invert
/flatten
/sum
/multiply
```

Any other path will yield an error.

As per specification, only square matrices with integers are permitted.

## Test

1. To run tests simply run:

```
go test ./pkg/... -cover
```

Example csv files are present in the folder `./test_data`.

---
