package search

import (
	"fmt"
	"log"
)

//Result保存搜索结果
type Result struct {
	Field   string
	Content string
}

//Matcher定义了要实现的
//新搜索类型的行为
type Matcher interface{
	Search(feed *Feed, searchTrem string) ([]*Result, error)
}

//Match函数，为每个数据源单独启动goroutine来执行这个函数
//并发的执行搜索
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	//使用特定的匹配器进行搜索
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	//将结果写入通道
	for _, result := range searchResults {
		results <- result
	}
}

//Display从每个单独的goroutine接收到结果之后
//在终端窗口输出
func Display(results chan *Result) {
	//通道会一直阻塞，直到有结果输入
	//一旦通道被关闭，for循环就会被终止
	for result := range results {
		fmt.Printf("%s : \n %s \n\n", result.Field, result.Content)
	}
}
