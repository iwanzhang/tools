package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/nsqio/go-nsq"
)

// NSQ Producer Demo

var producer *nsq.Producer

// 初始化生产者
func initProducer(str string) (err error) {
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer(str, config)
	if err != nil {
		fmt.Printf("create producer failed, err:%v\n", err)
		return err
	}
	return nil
}

func main() {
	//nsqAddress := "127.0.0.1:4150"
	nsqAddress := "127.0.0.1:4250"
	err := initProducer(nsqAddress)
	if err != nil {
		fmt.Printf("init producer failed, err:%v\n", err)
		return
	}

	//reader := bufio.NewReader(os.Stdin) // 从标准输入读取
	//for {
	//	data, err := reader.ReadString('\n')
	//	if err != nil {
	//		fmt.Printf("read string from stdin failed, err:%v\n", err)
	//		continue
	//	}
	//	data = strings.TrimSpace(data)
	//	if strings.ToUpper(data) == "Q" { // 输入Q退出
	//		break
	//	}
	//	// 向 'topic_demo' publish 数据
	//	err = producer.Publish("topic_demo", []byte(data))
	//	if err != nil {
	//		fmt.Printf("publish msg to nsq failed, err:%v\n", err)
	//		continue
	//	}
	//}
	//rand.Seed(time.Now().Unix())
	var data uint64
	for {
		data++
		_ = producer.Publish("topic_demo", []byte(strconv.FormatUint(data, 10)))
		time.Sleep(100 * time.Millisecond)
	}
}
