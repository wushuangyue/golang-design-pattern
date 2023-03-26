/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-03-26 20:46:40
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-03-26 23:43:44
 * @FilePath: \golang-design-pattern\1_CreationalPatterns\1_01_singleton\singleton_lazy.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package singleton

import (
	"fmt"
	"sync"
)

var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

// GetLazyInstance 懒汉式
func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
			fmt.Println("Create Singleton")
		})
	}
	return lazySingleton
}
