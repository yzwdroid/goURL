# Contributing to goURL

## Welcome!

This document will help you to get started.

## Setup

```bash
git clone https://github.com/yzwdroid/goURL.git
```
## Formatter

The gofmt tool rewrites code into the standard format.

```
gofmt filenames
```

## Linter

Linter is a tool that analyzes source code.

```	
go get -u golang.org/x/lint/golint
golint filenames
```

## Vim setup

```bash
augroup go
  autocmd!

  autocmd FileType go nmap <silent> <Leader>v <Plug>(go-def-vertical)
  autocmd FileType go nmap <silent> <Leader>s <Plug>(go-def-split)
  autocmd FileType go nmap <silent> <Leader>d <Plug>(go-def-tab)

  autocmd FileType go nmap <silent> <Leader>x <Plug>(go-doc-vertical)

  autocmd FileType go nmap <silent> <Leader>i <Plug>(go-info)
  autocmd FileType go nmap <silent> <Leader>l <Plug>(go-metalinter)

  autocmd FileType go nmap <silent> <leader>b :<C-u>call <SID>build_go_files()<CR>
  autocmd FileType go nmap <silent> <leader>t  <Plug>(go-test)
  autocmd FileType go nmap <silent> <leader>r  <Plug>(go-run)
  autocmd FileType go nmap <silent> <leader>e  <Plug>(go-install)

  autocmd FileType go nmap <silent> <Leader>c <Plug>(go-coverage-toggle)


  autocmd Filetype go command! -bang A call go#alternate#Switch(<bang>0, 'edit')
  autocmd Filetype go command! -bang AV call go#alternate#Switch(<bang>0, 'vsplit')
  autocmd Filetype go command! -bang AS call go#alternate#Switch(<bang>0, 'split')
  autocmd Filetype go command! -bang AT call go#alternate#Switch(<bang>0, 'tabe')
augroup END
```

## VSCode setup


Install the Go extension. You should immediately see a prompt in the bottom-right corner of your screen titled Analysis Tools Missing. This extension relies on a suite of command-line tools, which must be installed separately. Accept the prompt, or use the Go: Install/Update Tools command.

Also, see the tutorial [https://code.visualstudio.com/docs/languages/go](https://code.visualstudio.com/docs/languages/go)
