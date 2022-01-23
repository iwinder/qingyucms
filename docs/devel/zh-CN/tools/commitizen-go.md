# commitizen-go 使用

用于进入交互模式，并根据提示生成 Commit Message，然后提交

## 安装
### 方式一：全局

```shell
#  获取源文件，可直接下载对应1.0.0的zip包
git clone https://github.com/lintingzhen/commitizen-go.git
# 解压缩后，进入commitizen-go文件夹 
make && make install
# 使用
git cz 
```
### 方式二-不推荐
这种方式需要看被安装到哪个目录中，不然会提示找不到命令
```shell
go  install github.com/lintingzhen/commitizen-go@latest
# 使用，如果当前项目中有vendor文件夹，将会被安装到这里面，之后执行如下命令即可
commitizen-go
```
 

1.17开始，使用 `go get`将提示如下信息：
> go get: installing executables with 'go get' in module mode is deprecated.
To adjust and download dependencies of the current module, use 'go get -d'.
To install using requirements of the current module, use 'go install'.
To install ignoring the current module, use 'go install' with a version,
like 'go install example.com/cmd@latest'.
For more information, see https://golang.org/doc/go-get-install-deprecation
or run 'go help get' or 'go help install'.