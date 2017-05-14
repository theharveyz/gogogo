package main

import "fmt"

func worker(id int, jobs <-chan []int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started")
		go func() {
			for _, v := range j {
				fmt.Println(v)
				results <- v
			}
		}()

	}
}

func main() {
	/**
	 * 官网例子里不是很形象, 这里举个分片计算的例子
	 */
	ss := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}

	jobs := make(chan []int, 100)
	results := make(chan int, 100)

	var id int
	for _, w := range ss {
		id = id + 1
		go worker(id, jobs, results)
		fmt.Println(w)
		jobs <- w
	}
	close(jobs)

	f := 1
	for i := 0; i < 6; i++ {
		f *= <-results
	}
	fmt.Println("factorial result:", f)

}
