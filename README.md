# ip2region
### 用途
Go语言写的通过IP快速查看地区，准确率99.99999%，无需联网
### 安装
##### 1.根据系统下载编译好的最新releases包直接运行
https://github.com/lyy910203/ip2region/releases
##### 2.go编译 go>1.13
```
git clone https://github.com/lyy910203/ip2region.git
go build 
```

### 使用
##### 1.启动
linux:
```
chmod +x ip2region
ip2region --help
Usage of ./ip2region:
  -h string
    	run in host ip,default 0.0.0.0 (default "0.0.0.0")
  -p int
    	run in port,default 80 (default 80)

```
windows:
```
ip2region.exe --help
Usage of ./ip2region:
  -h string
    	run in host ip,default 0.0.0.0 (default "0.0.0.0")
  -p int
    	run in port,default 80 (default 80)

```
##### 2.查询
查询访问者的IP地区:访问http://{服务器IP}:{-p指定的端口，默认80}/ip
查询任意IP地区:访问http://{服务器IP}:{-p指定的端口，默认80}/ip?ip={待查询的IP}