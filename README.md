[![Build Status](https://travis-ci.org/agupta666/elf.svg?branch=master)](https://travis-ci.org/agupta666/elf)
# Elf

Elf is an HTTP server which lets you interact with it through a command-line interface. It gives you a quick and easy interface to set up `routes` and attach `actions` to them.

Elf provides a large collection of built-in actions, which can be added and configured with a fair amount of ease, from the command-line interface. And all of these can be achieved without writing a bit of code. The elf was primarily created for the purpose of quickly mocking HTTP based services without using going through the hassle of installing another scripting language.

Elf is a single executable that you can download and put in your path, to get started right away. Elf attempts to make it trivially easy to setup HTTP endpoints. The `elf` CLI provides a bash like command-line editing interface and intelligent `TAB` completions, that lets you manage routes and actions with ease.

> **NOTE:** Elf is an experiment and is expected to be used for testing purposes only. Elf is not intended for production usage.

## Install

```
$ go get github.com/agupta666/elf
```

## Quickstart

To start `elf`, fire the below command in a terminal. This starts Elf in interactive mode and drops
you to a prompt. Now you are ready to interact with the HTTP server listening on port 8080.

```
$ elf
starting default http endpoint 0.0.0.0:8080
elf>  

```
Hitting the `TAB` key will display a list of commands supported by Elf. See the command reference section for
details of the individual commands.

## Routes and Actions

Elf lets you set up routes using the `route` command, which is of the form

```
elf> route </some/path> <action>
```
Thus when a request is received that matches the given path, elf executes the `action` attached to it.

For example, the below route command will attach a built-in action [which will respond with the contents of the `data.txt` file] for a request matching the given path.

```
elf> route /data @data.txt
```
Thus sending a HTTP request using curl, displays the response shown below.

```
$ curl -i http://localhost:8080/data
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Date: Mon, 04 Sep 2017 18:40:49 GMT
Transfer-Encoding: chunked

contents of data.txt ...
```

## Built-in actions
| Action     | Syntax                                   | Meaning                                                               |
|------------|------------------------------------------|-----------------------------------------------------------------------|
|  File      |  @<file-path>                            | Responds with the contents of the given file                          |
|  Shell     |  !<shell-command>                        | Responds with the output of the shell command                         |
|  Markdown  |  #<path-to-markdown-file>                | Responds with HTML representation of the given markdown file          |
|  Redirect  |  ^<some-url>                             | Redirects to the given URL                                            |
|  Forward   |  %<some-url>                             | Forwards the request to the given URL                                 |
|  Data      |  data[options...]                        | Responds with random data                                             |
|  Image     |  image[options...]                       | Responds with image data                                              |
|  JSON      |  json[data-set]                          | Responds with JSON representation of a data set                       |
|  Upload    |  upload[options]                         | Saves the uploaded file to a specified folder                         |
|  Directory |  dir[options]                            | Serves files from the specified folder                                |
|  Dump      |  dump[options]                           | Dumps the incoming HTTP request to a specified folder                 |
|  Echo      |  echo[options]                           | Echoes back the incoming HTTP request                                 |


## Action Reference

### File
This is a built-in action that serves a static file. A static file can be specified with its absolute path with `@` as a prefix.

E.g.:

```
route /static @/var/www/index.html
```

### Shell Commands
This is a builtin action which executes the specified command on the local server and responds with the output. This action can
be specified by prefixing `!` before the shell command

#### Examples


```
elf> route /clip !pbpaste
```

```
elf> route /go !"w3m -dump https://golang.org/doc/effective_go.html"
```

### Generate arbitrary data
```
elf> route /download "data[9024,.bin]"
elf> route /download "data[9024,.bin,trivia]"
```

### Generate Images

`route /some/route "image[width, height, color, format, name]"`

```
elf> route /download "image[400,400]"
elf> route /download "image[400,400, #E32171]"
elf> route /download "image[400,400, #E32171, .png]"
elf> route /download "image[400,400, #E32171, .png, gopher]"
```

### Respond with the JSON representation of a Key-Value set

Create a new Key-Value set

```
elf> kvset dset a:1 b:2 c:3
```

now add a route with a JSON action

```
elf> route /api json[dset]
```


### Upload files
add a route with upload action

```
elf> route /save "upload[/path/to/some/folder, picture]"
```

upload file with curl

```
$ curl --form picture=@somefile.png http://localhost:8080/save
```
