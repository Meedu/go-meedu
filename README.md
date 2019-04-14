
## MeEdu 插件安装工具

因为php在fpm跑composer命令稳定性太差，所以使用go提供api服务来安装composer的依赖，从而使得插件安装更加方便。

### 编译

window#386:  

```
$ docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp -e GOOS=windows -e GOARCH=386 golang:latest go build -v
```

window#64:  

```
$ docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp -e GOOS=windows -e GOARCH=amd64 golang:latest go build -v
```

Linux#386:  

```
$ docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp -e GOOS=linux -e GOARCH=386 golang:latest go build -v
```

Linux#64:  

```
$ docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp -e GOOS=linux -e GOARCH=amd64 golang:latest go build -v
```

Mac#64:  

```
$ docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp -e GOOS=darwin -e GOARCH=amd64 golang:latest go build -v
```

### Usage

```
./go-meedu -address=0.0.0.0:8089
```

> 建议配置 supervisor 等进程管理软件一起使用。

### 服务

对外提供 `/install` API，参数如下：

| 参数 | 类型 | 说明 |
| --- | --- | --- |
| `php` | `string` | php执行路径 | 
| `composer` | `string` | composer路径 |
| `action` | `string` | composer行为，es:require,remove |
| `pkg` | `string` | 包名，如：monolog/monolog=dev-master |
| `dir` | `string` | 命令执行路径，一般为meedu所在根目录 |
| `addons` | `string` | 需要安装依赖的插件，用于回调 |
| `notify` | `string` | 回调URL，用于通知meedu的插件依赖安装状态 |


**示例：**  

```
http://127.0.0.1:8089/install?php=php&composer=C:\Users\Administrator\Desktop\go-meedu\test\composer.phar&action=requires&pkg=monolog/monolog&dir=C:\Users\Administrator\Desktop\go-meedu\test
```
