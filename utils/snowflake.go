package utils

import (
	"errors"
	"sync"
	"time"
)

const (
	workerBits  uint8 = 10                      // 每台机器(节点)的ID位数 10位最大可以有2^10=1024个节点
	numberBits  uint8 = 12                      // 表示每个集群下的每个节点，1毫秒内可生成的id序号的二进制位数 即每毫秒可生成 2^12-1=4096个唯一ID
	workerMax   int64 = -1 ^ (-1 << workerBits) // 节点ID的最大值，用于防止溢出
	numberMax   int64 = -1 ^ (-1 << numberBits) // 同上，用来表示生成id序号的最大值
	timeShift   uint8 = workerBits + numberBits // 时间戳向左的偏移量
	workerShift uint8 = numberBits              // 节点ID向左的偏移量
	epoch       int64 = 1525705533000           // 这个是我在写epoch这个变量时的时间戳(毫秒)
)

type Snowflake struct {
	mu        sync.Mutex // 添加互斥锁 确保并发安全
	timestamp int64      // 记录时间戳
	workerId  int64      // 该节点的ID
	number    int64      // 当前毫秒已经生成的id序列号(从0开始累加) 1毫秒内最多生成4096个ID
}

func NewSnowflake(workerId int64) (*Snowflake, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("worker id error")
	}
	return &Snowflake{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

func (w *Snowflake) GetId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()

	now := time.Now().UnixNano() / 1e6
	if w.timestamp == now {
		w.number++
		if w.number > numberMax {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}

	id := int64((now-epoch)<<timeShift | (w.workerId << workerShift) | (w.number))
	return id
}
