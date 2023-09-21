## 01 初始测试
```
package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

```

You will notice that we're using %d as our format strings rather than %q. That's because we want it to print an integer rather than a string.

## 02 编写函数

```
package integers

// Add takes two integers and returns the sum of them
func Add(x, y int) int {
	return x + y
}

```

函数前的注释 将出现在 Go Doc 中

## 03 添加 example

在 `test.go`` 里 ，添加 `ExampleAdd` 函数，在测试和GOdoc种会被体现

```
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
```
