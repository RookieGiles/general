package snowflake

import (
	"errors"
	"sync"
	"time"
)

var (
	// Epoch 起始偏移时间戳(2023-01-01 00:00:00的毫秒时间戳)
	//该时间一定要小于第一个id生成的时间,且尽量大(影响算法最终的有效可用时间)
	//有效可用时间 = self::EPOCH_OFFSET + (-1 ^ (-1 << self::TIMESTAMP_BITS))
	Epoch int64 = 1672502400000
	// NodeBits 节点位的位长度,10bit可支持1024个机器节点
	NodeBits uint8 = 10
	// SequenceBits 计数序列号位数12个bit,每个节点1毫秒可支持每秒生成16(0~4095)个序列号
	SequenceBits uint8 = 12

	NodeMax     int64 = -1 ^ (-1 << NodeBits)     // 节点ID的最大值，用于防止溢出
	SequenceMax int64 = -1 ^ (-1 << SequenceBits) // 同上，用来表示生成id序号的最大值

	timeShift uint8 = NodeBits + SequenceBits // 时间戳向左的偏移量
	NodeShift uint8 = SequenceBits            // 节点ID向左的偏移量
)

// Node 定一个node工作节点所需要的基本参数
type Node struct {
	mu        sync.Mutex // 添加互斥锁 确保并发安全
	timestamp int64      // 记录时间戳
	node      int64      // 该节点的Id
	step      int64      // 当前毫秒已经生成的id序列号(从0开始累加) 1毫秒内最多生成4096个ID
}

// NewNode 实例化一个工作节点
func NewNode(NodeId int64) (*Node, error) {
	// 要先检测workerId是否在上面定义的范围内
	if NodeId < 0 || NodeId > NodeMax {
		return nil, errors.New("giles.wang/general/snowflake NodeId excess quantity")
	}
	// 生成一个新节点
	return &Node{
		timestamp: 0,
		node:      NodeId,
		step:      0,
	}, nil
}

// GetId 生成方法一定要挂载在某个Node实例下，这样逻辑会比较清晰 指定某个节点生成id
func (n *Node) GetId() int64 {
	// 获取id最关键的一点 加锁 加锁 加锁
	n.mu.Lock()
	defer n.mu.Unlock() // 生成完成后记得 解锁 解锁 解锁

	// 获取生成时的时间戳
	var reset bool
	now := time.Now().UnixNano() / 1e6 // 纳秒转毫秒
	if reset = !(n.timestamp == now); !reset {
		n.step++

		// 这里要判断，当前工作节点是否在1毫秒内已经生成numberMax个ID
		if n.step > SequenceMax {
			// 如果当前工作节点在1毫秒内生成的ID已经超过上限 需要等待1毫秒再继续生成
			for now <= n.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
			reset = true
		}
	}

	if reset {
		// 如果当前时间与工作节点上一次生成ID的时间不一致 则需要重置工作节点生成ID的序号
		n.step = 0
		n.timestamp = now // 将机器上一次生成ID的时间更新为当前时间
	}

	// 第一段 now - epoch 为该算法目前已经奔跑了xxx毫秒
	// 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
	ID := int64((now-Epoch)<<timeShift | (n.node << NodeShift) | (n.step))
	return ID
}
