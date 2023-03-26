/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-03-25 14:27:40
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-03-26 23:58:20
 * @FilePath: \golang-design-pattern\1-CreationalPatterns\1_01_singleton\singleton_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package singleton_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	singleton "github.com/wushuangyue/golang-design-pattern/1_CreationalPatterns/1_01_singleton"

)

func TestGetLazyInstance(t *testing.T) {
	assert.Equal(t, singleton.GetLazyInstance(), singleton.GetLazyInstance())

}

/*
*
测试并发call多次，看看instance是不是同一个
*/
func TestCallMultiTimes(t *testing.T) {
	wg := sync.WaitGroup{}

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			obj := singleton.GetLazyInstance()
			fmt.Printf("%p\n", obj)
			wg.Done()
		}()
	}
	wg.Wait() //试试看不要这个语句会输出什么？O(∩_∩)O哈哈~

}
