# sim-backend
## Intro
This is a management system<br>
😀
## Install
本系统所有的`api`都写在`sim-backend`中。

终端输入，下载最新版后端代码

`git clone https://github.com/cutety/sim-backend.git`

该项目依赖于Go Module，因此需要确保环境变量中`Go Module = on`。

```shell
# 进入项目根目录
$ cd sim-backend
$ mv config/config-dev.toml config/config.toml
# 编辑配置文件
$ vim config/config.toml
```

插入以下配置到config

```toml
# 以下配置使用viper读取，相关问题请查看viper文档 https://github.com/spf13/viper
[server]
ip = "127.0.0.1" # 地址
port = ":3000" # 端口
version = "latest" 
name = "sim"
debug = true  
mode = "debug" 
jwtKey="skk2jmm13z" # jwtKey
[extensions]
[extensions.mysql]
host = "127.0.0.1" # 数据库地址
port = 3306 #数据库端口
username = "root" # 数据库username
password = "xxx" # 数据库密码
db = "sim" # 数据库名
[extension.redis] # redis 相关配置
host = "127.0.0.1"
port = 6379
db = 0
```



```shell
# 运行项目，第一次运行会下载相关包
$ go run main.go
```

