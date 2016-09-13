# hash

handy tool for testing http clients ...

# Shell action usage

```
hash> route /clip !pbpaste
hash> route /go !"w3m -dump https://golang.org/doc/effective_go.html"
```

### Data action usage
```
hash> route /download "data[9024,.bin]"
hash> route /download "data[9024,.bin,trivia]"
```

### Image action usage

`route /some/route "image[width, height, color, format, name]"`

```
hash> route /download "image[400,400]"
hash> route /download "image[400,400, #E32171]"
hash> route /download "image[400,400, #E32171, .png]"
hash> route /download "image[400,400, #E32171, .png, gopher]"
```

### JSON action usage
create a new kv set

```
hash> kvset dset a=1 b=2 c=3
```

now add a route with a JSON action

```
hash> route /api json[dset]
```


### Upload action
add a route with upload action

```
hash> route /save "upload[/path/to/some/folder, picture]"
```

upload file with curl

```
$ curl --form picture=@somefile.png http://localhost:8080/save
```
