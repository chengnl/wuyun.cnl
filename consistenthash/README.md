# 一致性hash
  一致性hash算法处理

# 功能说明
提供一致性hash获取逻辑。

# 使用说明

## 1、创建节点
根据自己的实际情况，创建节点信息，初始化一致hash函数
```
nodes := []interface{{},{},{}}
chash := consistenthash.NewConsistenHash(vNum, nodes)

```
如果需要自定义hash函数
```
nodes := []interface{{},{},{}}
chash := consistenthash.NewConsistenHash(vNum, func(hashString string) int{
   //自定义hash方法
},nodes)

```

## 2、参数说明

vNum:虚拟节点个数

hfunc:hash函数   默认使用FNV1A_32_HASH算法

nodes:使用对象节点  支持int，string，Node接口定义类型和自定义类型

## 3、使用
添加节点：
```
consistenthash.AddNode(节点参数)
```
删除节点：
```
consistenthash.DeleteNode(节点参数)
```
获取节点：
```
consistenthash.GetNode(内容参数)
```

# 完整示例
参考chash_test.go的示例

使用：

先下载：go get github.com/chengnl/wuyun.cnl/consistenthash/

使用引入：import "github.com/chengnl/wuyun.cnl/consistenthash"

注：
使用数据结构redblacktree：[https://github.com/emirpasic/gods/tree/master/trees/redblacktree](https://github.com/emirpasic/gods/tree/master/trees/redblacktree)

参看文档一致hash：[http://blog.csdn.net/cywosp/article/details/23397179/](http://blog.csdn.net/cywosp/article/details/23397179/)