# goURL

check dead links inside files.

![Alt Text](https://dev-to-uploads.s3.amazonaws.com/i/yczaka2e9vazsg49i4l8.gif)

# Usage

After clone the repo, 

```bash
export CLICOLOR=1

go build .

./goURL -f urls.txt

./goURL -f urls2.txt

// to ignore url patterns in file ignore.txt
./goURL -f urls2.txt -i true
```

# Features

- show usage
 
`./goURL`

- show version

`./goURL -v`

`./goURL --version`

- colorfull outputs
bad URLs in red, good URLs in green, unknown URLs in gray, redirect URLS in yellow,

- support parallelization using go routine

- optimize network code by only requesting for headers

