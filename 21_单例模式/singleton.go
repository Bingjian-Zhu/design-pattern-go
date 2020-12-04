package singleton

import "sync"

//Singleton 单例模式
type Singleton struct{}

var singleton *Singleton
var once sync.Once

//GetInstance 获取单例模式对象，双重检查加锁保证线程安全
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = new(Singleton)
	})

	return singleton
}
