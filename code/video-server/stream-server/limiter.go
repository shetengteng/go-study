package main

import "log"

// 流控模块：用户发起请求，大量请求时，server的连接数有限，带宽有限，导致系统不可用，如内存，带宽被用完
// 使用token bucket算法，令牌桶
// 一次request 得到一个token，当业务处理结束，返回reponse的时候，返还一个token
// token的个数限制了请求个数
// 使用chan进行共享线程的信息
// 使用chan实现token bucket算法

type ConnLimiter struct {
	concurrentNum int
	bucket        chan int
}

func CreateLimiter(num int) *ConnLimiter {
	limiter := ConnLimiter{}
	limiter.concurrentNum = num
	limiter.bucket = make(chan int, num) // buffer chan
	return &limiter
}

func (limiter *ConnLimiter) RequireToken() bool {
	if len(limiter.bucket) >= limiter.concurrentNum {
		// 超过限制，直接返回false
		log.Printf("reached the rate limiter")
		return false
	}
	// 说明有一个请求到来
	limiter.bucket <- 1
	return true
}

func (limiter *ConnLimiter) ReleaseToken() {
	// 释放资源
	c := <-limiter.bucket
	log.Printf("new connection is finish :%d", c)
}
