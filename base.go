package random

import (
	"math/rand"
	"time"
)

// 开始时间为 北京时间 2016/01/01 0:0:0
const StartUnixTime = 1462032000

var (
	rnd    *rand.Rand
	ch     chan int64
)

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	ch = make(chan int64, 1000)
	go randomBase(ch)
}

func randomBase(c chan int64) {
	for {
		c <- rnd.Int63()
	}
}
