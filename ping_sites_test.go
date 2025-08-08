package notes

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func for_test_parametrized(limit, timeout int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	sites := []string{"123", "456", "789"}
	ping_3_sites(ctx, sites, limit, 10, cancel)
	<-ctx.Done()
	fmt.Println("Done")
}

func TestPingSites(t *testing.T) {
	for_test_parametrized(100, 3)
	for_test_parametrized(1000, 3)
	for_test_parametrized(1000, 8)

}
