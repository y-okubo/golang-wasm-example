# golang-wasm-example

Example app using Go's wasm

## Build

```
$ GOOS=js GOARCH=wasm go build -o test.wasm test.go
```

## Run

```
$ go run server.go
```

## Access

http://localhost:5555/

## License

MIT

## Author

Yasuhrio Matsumoto (a.k.a. mattn)

## Update

```
curl -sO https://raw.githubusercontent.com/golang/go/master/misc/wasm/wasm_exec.html
curl -sO https://raw.githubusercontent.com/golang/go/master/misc/wasm/wasm_exec.js
```

## References

https://qiita.com/cia_rana/items/bbb4112b480636ab9d87