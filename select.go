// Go の _select_ を使うと、複数のチャネルで受信を待てる。
// ゴルーチンとチャネル、そして select の組み合わせは Go の強力な機能である。

package main

import (
	"fmt"
	"time"
)

func main() {

	// この例では2つのチャネルに select を使う。
	c1 := make(chan string)
	c2 := make(chan string)

	// いずれのチャネルも一定の時間が経てば値を受信する。
	// これはゴルーチンが RPC を同期的に実行する状況を模している。
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// `select` を使ってふたつの値をいずれも非同期で受信し、届いた方から表示する。
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	a1 := make(chan string)
	a2 := make(chan string)
	a3 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		a1 <- "Let's"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		a2 <- "learn"
	}()
	go func() {
		time.Sleep(3 * time.Second)
		a3 <- "Go!"
	}()

	for i := 0; i < 3; i++ {
		select {
		case step1 := <-a1:
			fmt.Println(step1)
		case step2 := <-a2:
			fmt.Println(step2)
		case step3 := <-a3:
			fmt.Println(step3)
		}
	}
}
