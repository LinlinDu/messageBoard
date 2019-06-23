# messageBoard
a restful message api

## installation
```
$go get github.com/LinlinDu/messageBoard
```

## How to run

### Required

- Mysql

### Ready

Create a **message_board database** and import [SQL](https://github.com/LinlinDu/messageBoard/blob/master/docs/message_board.sql)

### Conf

You should modify `conf/app.ini`

```
#debug or release
RunMode = release

[app]
JWTSecret = 23347$040412

ImagePrefixUrl = http://127.0.0.1:8000/
ImageSavePath = runtime/upload/images/
# MB
ImageMaxSize = 5
ImageAllowExts = .jpg,.jpeg,.png

LogSavePath = runtime/logs/
LogFileExt = log
TimeFormat = 20060102

[server]
HTTPPort = 8000
ReadTimeout = 60
WriteTimeout = 60
...
```

### Run
```
$cd $GOPATH/src/messageBoard
```
use port value in ini file 
```
$go run main.go
```
or set new port
```
$go run main.go -port 8080
```
## Features
+ RESTful API
+ log
+ Jwt-go
+ Gin
+ Graceful stop
+ App configurable
