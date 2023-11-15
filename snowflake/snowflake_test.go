package snowflake

import (
	"fmt"
	"testing"
)

func TestSnowFlake(t *testing.T) {
	node, err := NewNode(1)

	if err != nil {
		fmt.Println(err)
		return
	}
	//生成一个雪花Id
	snowflakeId := node.GetId()

	fmt.Println(snowflakeId)

	//并发生成雪花Id
	ch := make(chan int64)
	count := 10000
	open := make(chan struct{}, 1)
	for i := 0; i <= count; i++ {
		go func(i int) {
			if i == count {
				close(open)
			}
			<-open
			id := node.GetId()
			ch <- id
		}(i)
	}

	defer close(ch)

	m := make(map[int64]int)
	for i := 0; i <= count; i++ {
		id := <-ch
		_, ok := m[id]
		if ok {
			t.Error("Id is not unique\n")
			return
		}

		m[id] = i
	}

	fmt.Println("all", count, "snowflake ID Get successed!")
}
