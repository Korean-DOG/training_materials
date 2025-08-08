package notes

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Counters struct {
	counter map[string]int
	mutex   sync.Mutex
	limit   int
}

func (c *Counters) increase(name string) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	val, ok := c.counter[name]
	if !ok {
		c.counter[name] = 1
		return true
	} else if val == c.limit {
		return false
	} else {
		c.counter[name] += 1
		return true
	}
}

func (c *Counters) setLimit(new_limit int) { c.limit = new_limit }

func (c *Counters) notFinished() []string {
	var ret []string
	for k, v := range c.counter {
		if v < c.limit {
			ret = append(ret, k)
		}
	}
	return ret
}

func real_ping(addr string) {
	r := rand.Intn(100)
	time.Sleep(time.Duration(r) * time.Millisecond)
	//fmt.Println("Site reached: '", addr, "' in", r, "ms")
}

func ping(limit int) func(context.Context, string) bool {
	var c Counters
	c.counter = make(map[string]int)
	c.setLimit(limit)

	return func(ctx context.Context, addr string) bool {
		if addr == "" {
			fmt.Println(c.counter)
		}
		select {
		case <-ctx.Done():
			return false
		default:
			ok := c.increase(addr)
			if ok {
				real_ping(addr)
				return true
			} else if len(c.notFinished()) == 0 {
				return false
			}
			return ok
		}
	}
}

func worker(ctx context.Context, sites []string, ping_func func(context.Context, string) bool, cancel func()) {
	var sites_to_ping []string

	for i := 0; i < len(sites); i++ {
		sites_to_ping = append(sites_to_ping, sites[i])
	}

	for {
		if len(sites_to_ping) == 0 {
			cancel()
			return
		}
		for pos, site := range sites_to_ping {
			can_continue := ping_func(ctx, site)
			if !can_continue {
				if (len(sites_to_ping) - 1) > pos+1 {
					sites_to_ping = append(sites_to_ping[:pos], sites_to_ping[pos+1:]...)
				} else {
					sites_to_ping = sites_to_ping[:pos]
				}
			}
		}
	}
}

func ping_3_sites(ctx context.Context, sites []string, num_of_calls, limit_of_parallel int, cancel func()) {
	ping_func := ping(num_of_calls)

	for i := 0; i < limit_of_parallel; i++ {
		go worker(ctx, sites, ping_func, cancel)
	}

	defer func() { ping_func(ctx, "") }()

	started := time.Now()
	fmt.Println("started", started.Second())

	defer func() { fmt.Println("finish-defer", time.Now().Second()) }()

	for {
		select {
		case <-ctx.Done():
			switch ctx.Err() {
			case context.DeadlineExceeded:
				fmt.Println("Завершение по таймауту")
			case context.Canceled:
				fmt.Println("Завершение по отмене контекста")
			}
			done := time.Now()
			fmt.Println("break", done.Second())
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("Прошла ещё 1 секунда")
		}
	}
	fmt.Println("end", time.Now().Second())
}
