package eg1

import "testing"

func TestAll(t *testing.T) {
	var sub Subject = NewProxy()

	res := sub.Do()

	if res != "pre:real:after" {
		t.Fail()
	}
}
