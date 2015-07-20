# amobee-uploader

## Compiling

Assuming Go is setup properly (we're using the "recommended layout" (http://golang.org/doc/code.html)).

```
$> cd $GOPATH
$> mkdir -p github.com/kontera-technologies
$> cd github.com/kontera-technologies
$> git clone ...
$> cd amobee-uploader
$> go run main.go --aws-access-key-id=XXX --aws-secret-access-key=XXX --local-path=XXX --s3-path=XXX
```
