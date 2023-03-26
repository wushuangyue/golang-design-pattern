/*
 * @Author: wushuangyue wushyue@163.com
 * @Date: 2023-03-25 14:27:40
 * @LastEditors: wushuangyue wushyue@163.com
 * @LastEditTime: 2023-03-26 20:24:00
 * @FilePath: \golang-design-pattern\1-CreationalPatterns\1_01_singleton\singleton_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package singleton_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	singleton "github.com/wushuangyue/golang-design-pattern/1_CreationalPatterns/1_01_singleton"
)

func TestGetInstance(t *testing.T) {
	assert.Equal(t, singleton.GetInstance(), singleton.GetInstance())

}
