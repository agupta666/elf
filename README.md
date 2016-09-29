# wish

handy tool for testing http clients ...

# Shell action usage

```
wish> route /clip !pbpaste
wish> route /go !"w3m -dump https://golang.org/doc/effective_go.html"
```

### Data action usage
```
wish> route /download "data[9024,.bin]"
wish> route /download "data[9024,.bin,trivia]"
```

### Image action usage

`route /some/route "image[width, height, color, format, name]"`

```
wish> route /download "image[400,400]"
wish> route /download "image[400,400, #E32171]"
wish> route /download "image[400,400, #E32171, .png]"
wish> route /download "image[400,400, #E32171, .png, gopher]"
```

### JSON action usage
create a new kv set

```
wish> kvset dset a=1 b=2 c=3
```

now add a route with a JSON action

```
wish> route /api json[dset]
```


### Upload action
add a route with upload action

```
wish> route /save "upload[/path/to/some/folder, picture]"
```

upload file with curl

```
$ curl --form picture=@somefile.png http://localhost:8080/save
```
