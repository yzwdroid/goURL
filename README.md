# goURL

check dead links inside a file.

# Usage

After clone the repo, `cd goURL`
`go run main.go *.txt`
`go run main.go urls2.txt`

# Features

- Check usage
`go run main.go`

- Check version
`go run main.go -v`
`go run main.go --version`

- glob patterns
`go run main.go *.txt`

- colorfull output
bad URLs in red, good URLs in green, unknown URLs in gray

- support parallelization using go routine

- optimize network code by only request header

