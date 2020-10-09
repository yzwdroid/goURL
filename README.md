# goURL

check dead links inside files.

![Alt Text](https://dev-to-uploads.s3.amazonaws.com/i/yczaka2e9vazsg49i4l8.gif)

# Usage

After clone the repo, 

```bash
cd goURL

# For Windows
./goURL.exe urls.txt

# For Linux
./goURL-Linux urls.txt

# For macOS
./goURL-macOS urls.txt

OR

go run main.go *.txt

go run main.go urls2.txt
```

# Features

- Check usage
 
`go run main.go`

- Check version

`go run main.go -v`

`go run main.go --version`

- glob patterns

`go run main.go *.txt`

- colorfull outputs
bad URLs in red, good URLs in green, unknown URLs in gray, redirect URLS in yellow,

- support parallelization using go routine

- optimize network code by only requesting for headers

