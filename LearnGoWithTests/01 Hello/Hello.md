# Hello

## 00 Hello world

```
package main

import "fmt"

func main() {
	fmt.Print("Hello")
}

```

## 01 里外分离

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