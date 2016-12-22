#snowflake算法

go语言snowflake算法实现

# 使用说明

```
s, err := snowflake.Snowflake(1, 1)
s.GetSeqID()
```

# 参数说明
dataCenterID :数据中心ID

workID:工作机器ID

代码中使用spinlock自旋锁，在多协程并发的情况下获取序列比互斥锁速度快。

# 完整示例
参考snowflake_test.go的示例

使用：

先下载：go get github.com/chengnl/wuyun.cnl/snowflake/

使用引入：import "github.com/chengnl/wuyun.cnl/snowflake"

注：
参看文档snowflake[https://github.com/twitter/snowflake](https://github.com/twitter/snowflake)

参看文档[http://blog.csdn.net/yangding_/article/details/52768906](http://blog.csdn.net/yangding_/article/details/52768906)


