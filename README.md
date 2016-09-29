# elf

handy tool for testing http clients. Add routes and attach builtin actions to them and start testing your clients.

## Install

```
$ go get github.com/agupta666/elf
```

## Usage

### Shell action usage

```
elf> route /clip !pbpaste
elf> route /go !"w3m -dump https://golang.org/doc/effective_go.html"
```

### Data action usage
```
elf> route /download "data[9024,.bin]"
elf> route /download "data[9024,.bin,trivia]"
```

### Image action usage

`route /some/route "image[width, height, color, format, name]"`

```
elf> route /download "image[400,400]"
elf> route /download "image[400,400, #E32171]"
elf> route /download "image[400,400, #E32171, .png]"
elf> route /download "image[400,400, #E32171, .png, gopher]"
```

### JSON action usage
create a new kv set

```
elf> kvset dset a=1 b=2 c=3
```

now add a route with a JSON action

```
elf> route /api json[dset]
```


### Upload action
add a route with upload action

```
elf> route /save "upload[/path/to/some/folder, picture]"
```

upload file with curl

```
$ curl --form picture=@somefile.png http://localhost:8080/save
```
