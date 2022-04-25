package main

import "time"

type Limiter struct {
	Capacity int
	Interval time.Duration
	AccessNums  int
	LastAccessTime time.Time
}


func (limiter *Limiter) AccessControl() bool{
	tNow := time.Now()
	diffTime := tNow.Sub(limiter.LastAccessTime)

	// 时间窗口内 允许多少的访问量
	leaks := int(float64(diffTime) / float64(limiter.Interval))
	if leaks > 0{
		if limiter.AccessNums <= leaks {
			limiter.AccessNums = 0
		}else{
			limiter.AccessNums -= leaks
		}
		limiter.LastAccessTime = tNow
	}

	if limiter.Capacity >=  limiter.AccessNums{
		return true
	}
	return false
}

