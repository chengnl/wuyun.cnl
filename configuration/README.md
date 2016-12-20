# 读取配置文件
  配置文件读取处理

# 功能说明
提供配置文件读取处理

# 使用说明

## 1、初始化
加载配置文件
```
configuration.InitConfigFile(filePath)

```

提供文件路径加载配置文件

## 2、使用
提供获取配置文件key值，支持返回string，bool，float，int

其中float支持float32，float64

int支持int，int8,int16,int32,int64；uint8,uint16,uint32,uint64;

另外支持获取时缺省值设置。

例如获取int8：

```
intVal, intErr := configuration.GetInt("int8", 10, 8)

```
缺省值获取方式：

```
intVal, intErr := configuration.GetIntDefaultVal("int8", 10, 8,56)

```


# 完整示例
参考config_test.go的示例

使用：

先下载：go get github.com/chengnl/wuyun.cnl/configuration/

使用引入：import "github.com/chengnl/wuyun.cnl/configuration"

#配置文件格式

注释使用#或者!开头，例如：
```
#键值对说明
```
键值对使用=或者:赋值，例如：
```
key=value
```