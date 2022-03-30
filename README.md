# fix-resolveEmbed-on-root
This repo. is intended to show a result of fixing Go's `embed` function in a specific situation.

## License
This repo. contains a part of Go's source codes to build fixed binary. The codes are under the original Go's license (show ./src/LICENSE). The other parts of this repo. is under MIT license.

## Usage
```shell
git clone https://github.com/tenkoh/fix-resolveEmbed-on-root
cd fix-resolveEmbed-on-root
docker build -t {image name} .
docker run -it --name {container name} {image name}
```

## What does this commit improve?
With the latest go(1.18), this command fails.
```shell
go build .
# /main.go:9:5: embed oo/foo.txt: open /oo/foo.txt: no such file or directory
```

This fail occurs only when the package directory is "/" (i.e. root) and the embed target is `embed.FS`. Please see the [issue #49570](https://github.com/golang/go/issues/49570) for detail. This commit fixs it.

```shell
# with built binary
/build/go/bin/go build .
# no error
```
