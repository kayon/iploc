# iploc

[![Build Status](https://travis-ci.org/kayon/iploc.svg?branch=master)](https://travis-ci.org/kayon/iploc)

```
go get github.com/kayon/iploc
```

使用纯真IP库 `qqwry.dat`，高性能，线程安全

> 附带的 `qqwry.dat` 为 `UTF-8` 编码 `2018-05-10版本`

## 更新 qqwry.dat

在纯真官网[下载最新的 qqwry.dat](http://www.cz88.net/fox/ipdat.shtml) 并转换为 `UTF-8` 使用命令行工具 [iploc-conv](#iploc-conv)


## Benchmarks

```
// 缓存索引
BenchmarkFind-8            	 2000000	       735 ns/op               136万/秒
// 无索引
BenchmarkFindUnIndexed-8   	   20000	     78221 ns/op               1.2万/秒
// 缓存索引，并发
BenchmarkFindParallel-8    	10000000	       205 ns/op               487万/秒
```

## 使用

```
func main() {
	loc, err := iploc.Open("qqwry.dat")
	if err != nil {
		panic(err)
	}
	detail := loc.Find("8.8.8") // 补全为8.8.0.8, 参考 ping 工具
	fmt.Printf("IP:%s; 网段:%s - %s; %s\n", detail.IP, detail.Start, detail.End, detail)
	
	detail2 := loc.Find("8.8.3.1")
	fmt.Printf("%t %t\n", detail.In(detail2.IP.String()), detail.String() == detail2.String())

	// output
	// IP:8.8.0.8; 网段: 8.7.245.0 - 8.8.3.255; 美国 科罗拉多州布隆菲尔德市Level 3通信股份有限公司
	// true true
}	
```

#### 快捷方法
#####Find(qqwrySrc, ip string) (*Detail, error)
`iploc.Find` 使用 `OpenWithoutIndexes`

#### 初始化
#####Open(qqwrySrc string) (*Locator, error)

`iploc.Open` 缓存并索引，生成索引需要耗时500毫秒左右，但会带来更高的查询性能

#####OpenWithoutIndexes(qqwrySrc string) (*Locator, error)

`iploc.OpenWithoutIndexes` 只读取文件头做简单检查，无索引

#### 查询

```
(*Locator).Find(ip string) *Detail

```
> 如果IP不合法，返回 `nil`


## 命令行工具

####iploc

命令行版IP查询

```
$ iploc 127.1
$ 127.0.0.1 本机地址 N/A
```
> DAT编译到二进制执行文件中，不依赖 `qqwry.dat` 位置

####<a name="iploc-conv"></a>iploc-conv

将原版 `qqwry.dat` 由 `GBK` 转换为 `UTF-8`

```
$ iploc-conv -s src.gbk.dat -d dst.utf8.dat
```

> 新生成的DAT文件保留原数据结构，由于编码原因，文件会增大一点

> 修正原 qqwry.dat 中几处错误的重定向 (qqwry.dat 2018-05-10)，并将 "CZ88.NET" 替换为 "N/A"

####iploc-gen

创建静态版本的 **iploc** 集成到你的项目中

`iploc-gen` 会在当前目录创建 iploc-binary.go 文件，拷贝到你的项目中，通过变量名 *IPLoc* 直接可以使用

```
$ iploc-gen /path/qqwry.dat
```

> `--pkg` 指定 package name, 默认 main

> `-n` 使用 `OpenWithoutIndexes` 初始化，无索引

## 静态编译 iploc 和 qqwry.dat 并集成到你的项目中

编译后的二进制没有 `qqwry.dat` 依赖，不需要再带着 `qqwry.dat` 一起打包了

#####示例

到项目目录 `$GOPATH/src/myproject/` 中

```
$ mkdir myloc && cd myloc
$ iploc-gen $GOPATH/src/github.com/kayon/iploc/qqwry.dat --pkg myloc
```

> $GOPATH/src/myproject/main.go

```
package main
	
import (
	"fmt"
	
	"myproject/myloc"
)
	
func main() {
	fmt.Println(myloc.IPLoc.Find("8.8.8.8"))
}
```