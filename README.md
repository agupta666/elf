# hash

handy tool for testing http clients ...

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
