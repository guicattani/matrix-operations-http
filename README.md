# League Backend Challenge

#### Submission by Guilherme Cattani

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

Example csv files are present in the folder `./test`.

---

## Challenge Specification

Given an uploaded csv file

```
1,2,3
4,5,6
7,8,9
```

1. Echo (given)

   - Return the matrix as a string in matrix format.

   ```
   // Expected output
   1,2,3
   4,5,6
   7,8,9
   ```

2. Invert
   - Return the matrix as a string in matrix format where the columns and rows are inverted
   ```
   // Expected output
   1,4,7
   2,5,8
   3,6,9
   ```
3. Flatten
   - Return the matrix as a 1 line string, with values separated by commas.
   ```
   // Expected output
   1,2,3,4,5,6,7,8,9
   ```
4. Sum
   - Return the sum of the integers in the matrix
   ```
   // Expected output
   45
   ```
5. Multiply
   - Return the product of the integers in the matrix
   ```
   // Expected output
   362880
   ```

The input file to these functions is a matrix, of any dimension where the number of rows are equal to the number of columns (square). Each value is an integer, and there is no header row. matrix.csv is example valid input.
