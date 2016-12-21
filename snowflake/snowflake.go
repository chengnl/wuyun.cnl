package snowflake

import (
	"errors"
	"fmt"
	"time"

	"github.com/chengnl/wuyun.cnl/spinlock"
)

const seqBit uint = 12                             //序列位数
const dataCenterBit uint = 5                       //数据中心占位
const workerBit uint = 5                           //工作机器占位
const timeBit uint = 41                            //时间戳占位
const maxSeqID int64 = 1<<seqBit - 1               //最大序列号值
const maxDataCenterID int64 = 1<<dataCenterBit - 1 //最大数据中心ID
const maxWorkerID int64 = 1<<workerBit - 1         //最大工作机器ID
const maxTimeID int64 = 1<<timeBit - 1             //最大时间戳ID
const timeShift = seqBit + dataCenterBit + workerBit
const dataCenterShift = seqBit + workerBit

var initTime = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC).UnixNano() / 1e6

type SnowFlake struct {
	seq           int64
	lastTimeStamp int64
	_dataCenterID int64
	_workID       int64
	//lock          sync.Mutex
	lock spinlock.SpinLock //使用自旋锁，并发获取效率高
}

func Snowflake(dataCenterID int64, workID int64) (*SnowFlake, error) {
	if dataCenterID > maxDataCenterID || dataCenterID < 0 {
		return nil, fmt.Errorf("dataCenterID can't be greater than： %d,or less than ：", maxDataCenterID, 0)
	}
	if workID > maxWorkerID || workID < 0 {
		return nil, fmt.Errorf("workID can't be greater than： %d，or less than ", maxWorkerID, 0)
	}
	return &SnowFlake{seq: 0, lastTimeStamp: -1, _dataCenterID: dataCenterID, _workID: workID}, nil
}
func (s *SnowFlake) GetSeqID() (int64, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	nowTime := s.genTimeStamp()
	ts := nowTime - s.lastTimeStamp
	if ts == 0 {
		s.seq = (s.seq + 1) & maxSeqID
		if s.seq == 0 {
			for nowTime <= s.lastTimeStamp {
				nowTime = s.genTimeStamp()
			}
		}
	} else if ts > 0 {
		s.seq = 0
	} else {
		return -1, errors.New("local time is error")
	}
	if nowTime > maxTimeID {
		return -1, errors.New("time is exceed")
	}
	s.lastTimeStamp = nowTime
	return nowTime<<timeShift | s._dataCenterID<<dataCenterShift | s._workID<<seqBit | s.seq, nil
}
func (s *SnowFlake) genTimeStamp() int64 {
	nowTime := time.Now().UnixNano() / 1e6
	return nowTime - initTime
}
