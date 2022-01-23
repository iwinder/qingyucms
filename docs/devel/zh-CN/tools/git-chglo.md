# git-chglo 使用指南

根据 Commit Message 生成 CHANGELOG

## 安装

```shell
# Go1.17开始提示go get: installing executables with 'go get' in module mode is deprecated.
go get -u github.com/git-chglog/git-chglog/cmd/git-chglog
#  方式二
go install github.com/git-chglog/git-chglog/cmd/git-chglog@latest
```

## 使用

```shell
git-chglog --init
```


选择
- ? What is the URL of your repository? https://gitee.com/windcoder/qingyublog
- ? What is your favorite style? github
- ? Choose the format of your favorite commit message <type>(<scope>): <subject>
- ? What is your favorite template style? standard
- ? Do you include Merge Commit in CHANGELOG? No
- ? Do you include Revert Commit in CHANGELOG? Yes
- ? In which directory do you output configuration files and templates? .chglog


```shell
# 前置准备，先有对应的tag
git tag -a 0.1.0  -m "Test Edition v0.1.0"
# 生成新日志
git-chglog -o CHANGELOG/CHANGELOG-0.1.md
```
