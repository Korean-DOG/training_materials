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

func ping(limit int) func(context.Context, string) (Counters, bool) {
	var c Counters
	c.counter = make(map[string]int)
	c.setLimit(limit)

	return func(ctx context.Context, addr string) (Counters, bool) {
		select {
		case <-ctx.Done():
			return c, false
		default:
			ok := c.increase(addr)
			if ok {
				real_ping(addr)
				return c, true
			} else if len(c.notFinished()) == 0 {
				return c, false
			}
			return c, ok
		}
	}
}

func worker(ctx context.Context, sites []string, ping_func func(context.Context, string) (Counters, bool), cancel func()) {
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
			_, can_continue := ping_func(ctx, site)
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

func ping_3_sites(ctx context.Context, sites []string, num_of_calls, limit_of_parallel int, cancel func()) (reason error, pinged map[string]int) {

	//reason = error("unknown")

	ping_func := ping(num_of_calls)

	for i := 0; i < limit_of_parallel; i++ {
		go worker(ctx, sites, ping_func, cancel)
	}

	// "return reason and counter"
	defer func() {
		pinged_counters, _ := ping_func(ctx, "")
		pinged = pinged_counters.counter
	}()

	started := time.Now()
	fmt.Println("started", started.Second())

	defer func() { fmt.Println("finish-defer", time.Now().Second()) }()

	for {
		select {
		case <-ctx.Done():
			switch ctx.Err() {
			case context.DeadlineExceeded:
				reason = context.DeadlineExceeded
			case context.Canceled:
				reason = context.Canceled
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
	return
}
