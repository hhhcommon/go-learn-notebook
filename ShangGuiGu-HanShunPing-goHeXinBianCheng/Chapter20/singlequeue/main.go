// 单行队列

// 单行队列的本质是一个有序数组遵循先进先出的原则。
// 这里面重点有两个：
// 1. 有序数组
// 2. 先进先出

// 也就是进去是这样进去的。 「0，1，2，3，4，5，6，7，8，9，10」
// 出来的顺序肯定是 0， 1，2，3，4，5，6，7，8，9，10.
// 队列的快照是可以根据不同的处理方式而不同的。
package main

import (
	"errors"
	"fmt"
	"os"
)

type SingleQueue struct {
	MaxSize int   // 队列容量
	Queue   [5]int // 队列数组
	Rear    int   // 指向队尾索引
	Front   int   // 指向队首索引
}

func (this *SingleQueue) AddSingleQueue(value int) (err error) {

	// 1. 判断队列是否已满
	if this.Rear == this.MaxSize-1 {
		return errors.New("Singlequeue  full!")
	}

	// 2. 队列没有满
	// 移动Rear
	this.Rear++
	this.Queue[this.Rear] = value
	return
}

func (this *SingleQueue) GetSingleQueue() (value int, err error) {

	// 1. 判断是否还有元素
	if this.Front == this.Rear {
		return -1, errors.New("Singlequeue empty!")
	}

	this.Front++
	value = this.Queue[this.Front]
	return
}

func (this *SingleQueue) ShowSingleQueue(){
	// 1.找到队首遍历到队尾, 不包含队首索引的元素
	for i := this.Front + 1; i < this.Rear+1; i++ {
		fmt.Printf("Queue : index = %d  value = %d\n", i, this.Queue[i])
	}
	fmt.Println("\n")

	fmt.Printf("SingleQueue： %p\n", this)
	fmt.Printf("maxsize： %v\n", this.MaxSize)
	fmt.Printf("queue: %v\n", this.Queue)
	fmt.Printf("front %v\n", this.Front)
	fmt.Printf("rear: %v\n", this.Rear)

	fmt.Println("\n")

}

func main() {

	// var queue SingleQueue

	// queue.MaxSize = 4
	// queue.Queue = -1
	// queue.Rear = -1

	queue := &SingleQueue{
		MaxSize: 5,
		Front: -1,
		Rear: -1,
	}

	var val int
	var key string

	for {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示显示队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.AddSingleQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {

				fmt.Println("加入队列ok")
			}
		case "get":
			val, err := queue.GetSingleQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "show":
			queue.ShowSingleQueue()
		case "exit":
			os.Exit(0)
		}

	}
}
