package eg2

import "testing"

func TestAll(t *testing.T) {
	api := NewFacadeAPI()
	ret := api.Test()
	var expect = "A module running\nB module running"
	if ret != expect {
		t.Fatalf("expect %s, return %s", expect, ret)
	}
}
