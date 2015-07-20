# amobee-uploader

## Compiling

Assuming Go is setup properly (we're using the "recommended layout" described here: http://golang.org/doc/code.html).

```
$> cd $GOPATH
$> mkdir -p github.com/kontera-technologies
$> cd github.com/kontera-technologies
$> git clone https://github.com/kontera-technologies/amobee-uploader
$> cd amobee-uploader
$> go run main.go --aws-access-key-id=XXX --aws-secret-access-key=XXX --local-path=XXX --s3-path=XXX
```
