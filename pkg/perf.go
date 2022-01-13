package pkgu

import "time"

type PerfChecker struct {
	start int64
	perf  int64
}

func NewPerf() PerfChecker {
	return PerfChecker{start: time.Now().UnixNano()}
}

func (p *PerfChecker) Ns() int64 {
	p.perf = (time.Now().UnixNano() - p.start) / 1000
	return p.perf
}
