package main

import "testing"

/**
测试程序需要在一个名为 xxx_test.go 的文件中编写
测试函数的命名必须以单词 Test 开始
测试函数只接受一个参数 t *testing.T
*/
func TestHello(t *testing.T) {
	s := Hello("Hello")
	w := "Hello World!"
	if s != w {
		t.Errorf("s=%s, w=%s", s, w)
	}
}

func TestHello2(t *testing.T) {

	//辅助函数
	helpMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	//子测试
	t.Run("sub test 1", func(t *testing.T) {
		s := Hello("Hello")
		w := "Hello World!"
		helpMessage(t, s, w)
	})

	//子测试
	t.Run("sub test 2", func(t *testing.T) {
		s := Hello("")
		w := "Hello World!"
		helpMessage(t, s, w)
	})
}
