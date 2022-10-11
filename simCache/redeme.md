##go面试题--实现一个内存缓存系统
####1.支持设定过期时间，精度到秒
####2.支持设定最大内存，在内存超出时做出合适的处理
####3.支持并发安全
####4.按照以下接口要求做出实现
```bash
   type Cache interface{
   	//size  1KB 100KB 1MB 2MB 1GB
   	SetMaxMemory(size string) bool
   	//将value 写入缓存
   	Set(key string,val interface{},expire time.Duration) bool
   	//根据key获取value
   	Get(key string) (interface{},error)
   	//删除key值
   	Del(key string) bool
   	//判断是否存在key
   	Exists(key string)bool
   	//清空所有的key
   	Flush() bool
   	//获取缓存中所有key的数量
   	Keys() int64
   }
```
####5.使用实例
    cache :=NewMemCache() 
    cache.SetMaxMemory("100MB")
    cache.Set("int",1)
    cache.Set("bool",false)
    cache.Set("data",map[string]interface{}{"a":1})
    cache.Get("int")
    cache.Del("int")
    cache.Flush()
    cache.Keys()

##加分项:
####6.处理过期时间数据进行删除
####7.代理模式
####8.单元测试


