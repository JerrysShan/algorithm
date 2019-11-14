package ratelimit

import "time"

// TokenBucket 令牌桶
type TokenBucket struct {
	capacity            int64
	lastRefillTimestamp int64
	refillPerSecond     int64
	availableToken      int64
	windowTimeInSeconds int64
}

// NewTokenBucket init TokenBucket instance
func NewTokenBucket(capacity, windowTimeInSeconds int64) *TokenBucket {
	return &TokenBucket{
		capacity:            capacity,
		lastRefillTimestamp: time.Now().Unix(),
		windowTimeInSeconds: windowTimeInSeconds,
		refillPerSecond:     capacity / windowTimeInSeconds,
		availableToken:      0,
	}
}

// TryAcquire try to acquire token
func (tb *TokenBucket) TryAcquire() bool {
	tb.refill()
	if tb.availableToken > 0 {
		tb.availableToken--
		return true
	}
	return false
}

func (tb *TokenBucket) refill() {
	now := time.Now().Unix()
	if now > tb.lastRefillTimestamp {
		elapsed := now - tb.lastRefillTimestamp
		tokensAdd := elapsed * tb.refillPerSecond
		tb.availableToken += tokensAdd
		if tb.availableToken > tb.capacity {
			tb.availableToken = tb.capacity
		}
	}
}
