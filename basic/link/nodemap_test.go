package link

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

//func TestLinkNode_get(t *testing.T) {
//	type fields struct {
//		Head *Node
//		Size int
//	}
//	type args struct {
//		key string
//		val int
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   int
//	}{
//		{
//			"TODO: Add test cases.",
//			fields{&Node{"", 0, nil}, 0},
//			args{"xx", 1},
//			0,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			ln := &LinkNode{
//				Head: tt.fields.Head,
//				Size: tt.fields.Size,
//			}
//			if got := ln.get(tt.args.key); got != tt.want {
//				t.Errorf("get() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestLinkNode_put(t *testing.T) {
	type fields struct {
		Head *Node
		Size int
	}
	type args struct {
		key string
		val int
	}
	type result struct{
		size int
		hasNext bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want result
	}{
		{
			"not first insert",
			fields{&Node{"yy", 1, nil}, 1},
			args{"xx", 1},
			result{2, true},
		},
		{
			"first insert",
			fields{&Node{"", 0, nil}, 0},
			args{"xx", 1},
			result{1, false},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ln := &LinkNode{
				Head: tt.fields.Head,
				Size: tt.fields.Size,
			}
			ln.put(tt.args.key, tt.args.val)
			//assert.Equal(t, ln.Head.Key, tt.args.key)
			//assert.Equal(t, ln.Head.Val, tt.args.val)
			assert.Equal(t, ln.Size, tt.want.size)
			assert.Equal(t, ln.Head.hasNext(), tt.want.hasNext)
		})
	}
}

//func TestNewNode(t *testing.T) {
//	tests := []struct {
//		name string
//		want *Node
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewNode(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewNode() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestNode_hasNext(t *testing.T) {
//	type fields struct {
//		Key  string
//		Val  int
//		Next *Node
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			n := &Node{
//				Key:  tt.fields.Key,
//				Val:  tt.fields.Val,
//				Next: tt.fields.Next,
//			}
//			if got := n.hasNext(); got != tt.want {
//				t.Errorf("hasNext() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}