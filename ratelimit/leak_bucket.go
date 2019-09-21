package ratelimit

import (
	"time"
)

// LeakBucket 漏桶法
type LeakBucket struct {
	rate         int64
	capacity     int64
	lastLeakTime int64
	water        int64
}

// NewLeakBucket create LeakBucket instance
func NewLeakBucket(rate, capacity int64) *LeakBucket {
	return &LeakBucket{
		rate,
		capacity,
		time.Now().Unix(),
		0,
	}
}

// Access 返回是否可访问
func (lk *LeakBucket) Access() bool {
	temp := (time.Now().Unix() - lk.lastLeakTime) * lk.rate
	lk.water -= temp
	if lk.water < 0 {
		lk.water = 0
	}
	lk.lastLeakTime = time.Now().Unix()
	if lk.water < lk.capacity {
		lk.water++
		return true
	}
	return false
}
