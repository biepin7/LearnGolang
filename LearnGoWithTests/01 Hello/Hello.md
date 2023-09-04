# Hello

## 00 Hello world

```
package main

import "fmt"

func main() {
	fmt.Print("Hello")
}

```

## 01 里外分离 方便测试

```
package main

import "fmt"

func main() {
	fmt.Print(Hello())
}

func Hello() string {
	return "Hello World"
}

```
## 针对 01 的测试

```
package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello World"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}

```

t.Errorf
我们调用 t 的 Errorf 方法打印一条消息并使测试失败。
f 表示格式化，它允许我们构建一个字符串，并将值插入占位符值 %q 中。当你的测试失败时，它能够让你清楚是什么错误导致的。

## 03 什么是测试驱动开发？

下一个需求是指定问候的接受者。
让我们从在测试中捕获这些需求开始。这是基本的测试驱动开发，可以确保我们的测试用例 真正 在测试我们想要的功能。
当你回顾编写的测试时，存在一个风险：即使代码没有按照预期工作，测试也可能继续通过。

## 04 真正的先有 test

```
package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello,Chris"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
```

## 05 由 04 对应的 hello

```
package main

func Hello(name string) string {
	return "Hello,"+name
}
```

## 06 优雅永不过时：加入常量和重构

```
package main

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name
}

```
