package cache

// Cache 缓存接口
type Cache interface {
    Set(key string, value interface{}) // 设置/添加一个缓存，如果 key 存在，用新值覆盖旧值
    Get(key string) interface{}        // 通过 key 获取一个缓存值
    Del(key string)                    // 通过 key 删除一个缓存值
    DelOldest()                        // 删除最"无用"的一个缓存值
    Len() int                          // 获取缓存已存在的记录数
}
