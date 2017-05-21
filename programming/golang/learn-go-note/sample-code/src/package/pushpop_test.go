package stack

import "testing"

func TestPushPop( t *testing.T ) {
	c := new(Stack)
	c.Push(5)
	res := c.Pop()
	if res != 5 {
		t.Log( "Pop doesn't give 5: ", res )
		t.Fail()
	}
}
