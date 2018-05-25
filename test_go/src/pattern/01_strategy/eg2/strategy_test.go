/*
*  策略模式:定义了一族算法，分别封装起来，让他们之间可以互相替换，此模式让算法的
*  			的变化独立于使用算法的客户。
*/
package eg2

import (
	"testing"
)

func TestAll(t *testing.T) {
	pc := NewPaymentContext("Ada", "", 123, &Cash{})
	pc.Pay()

	pc2 := NewPaymentContext("Bob", "0002", 888, &Bank{})
	pc2.Pay()
}
