# gounzip
go版本unzip实现，解决linux下zip文件解压乱码中文乱码问题，统一输出utf8编码。

使用方法:
项目引入

import	"github.com/Wang/gounzip/unzip"
unzip.Do(zipfile, targetDirectory)

可以参考 main.go实现
