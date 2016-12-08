# pool 对象池

go语言对象池

# 功能说明
提供对象池功能，包括对象获取，返回，销毁处理；其中获取提供超时等待设置。

# 使用说明

## 1、实现对象创建工厂
实现对象创建工厂PoolObjectFactoryer接口,该工厂实现以下接口方法：
```
//创建对象
CreateObj() (interface{}, error)
//销毁对象
DestroyObj(interface{}) error
//验证对象
ValidateObj(interface{}) error
```
## 2、对象池参数设置
提供连接池参数设置，参数设置如下：

```
//对象池配置信息
type PoolConfig struct {
	//最大空闲对象数量
	MAX_IDLE int
	//初始化最小空闲对象数量
	MIN_IDLE int
	//最大存活对象数量
	MAX_ACTIVE int
	//获取对象最大等待时间，<=0 一直等待
	MAX_WAIT time.Duration
	//是否在获取对象的时候验证对象有效性
	IsVaildObj bool
}
```
创建配置项示例：

```
config := &pool.PoolConfig{MAX_ACTIVE: 100, MAX_IDLE: 50, MIN_IDLE: 5, MAX_WAIT: 60 * time.Millisecond, IsVaildObj: true}
```

## 3、对象池创建
对象池创建，默认不给参数配置方式：
```
pool, err := pool.DefaultGenObjectPool(NewFactory())
```
此方式，默认MAX_ACTIVE和MAX_IDLE为5，MIN_IDLE为0，MAX_WAIT为-1一直等待对象返回，IsVaildObj为false，不检测对象有效性

提供参数配置方式：
```
config := &pool.PoolConfig{MAX_ACTIVE: 100, MAX_IDLE: 50, MIN_IDLE: 5, MAX_WAIT: 60 * time.Millisecond, IsVaildObj: true}
pool, err := pool.NewGenObjectPool(NewFactory(), config)
```

## 4、对象池使用
创建对象池后，根据需要预加载对象池：
```
pool.InitPool()
```
初始化默认创建MIN_IDLE对象放入对象池中

获取对象：
```
poolObject, err := pool.GetObject()
```
释放对象：
```
pool.ReturnObject(poolObject)
```
销毁对象：
```
pool.DestroyObject(poolObject)
```
清除对象池：
```
pool.Clear()
```
获取空闲对象数：
```
pool.GetNumIdle()
```
获取活动对象数：
```
pool.GetNumActive()
```
# 完整示例
参考example目录的示例
使用：
先下载：go get github.com/chengnl/wuyun.cnl/pool/

使用引入：import "github.com/chengnl/wuyun.cnl/pool"

注：
参考借鉴[https://github.com/fatih/pool](https://github.com/fatih/pool)，实现更加通用的对象池
