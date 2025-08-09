package notes

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func for_test_parametrized(t *testing.T, limit, timeout int, expected_reason error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	sites := []string{"123", "456", "789"}

	start_time := time.Now()
	reason, pinged := ping_3_sites(ctx, sites, limit, 10, cancel)
	spend_time := time.Since(start_time)

	<-ctx.Done()
	fmt.Println("*** Done for", spend_time, "seconds with reason", reason, "pinged:", pinged)

	test_name := fmt.Sprintf("Case w/limit=%d w/timeout=%d", limit, timeout)

	assert.Equal(t, reason, expected_reason, fmt.Sprintf("[%s] Reason %s is not as expected: %s", test_name, reason.Error(), expected_reason.Error()))

	if expected_reason == context.Canceled {
		for site, count := range pinged {
			assert.Equal(t, count, limit, fmt.Sprintf("[%s] Site '%s' pinged '%d' times that less than expected: '%d'", test_name, site, count, limit))
		}

		assert.Less(t, spend_time, time.Duration(timeout)*time.Second, fmt.Sprintf("[%s] Execution stops by reaching limit, so it must take less (%s) that timeout (%d)", test_name, spend_time.String(), timeout))
	} else {
		for site, count := range pinged {
			assert.NotEqual(t, count, 0, fmt.Sprintf("[%s] Site '%s' is not pinged'", test_name, site))
			assert.LessOrEqual(t, count, limit, fmt.Sprintf("[%s] Site '%s' pinged '%d' times that less than expected: '%d'", test_name, site, count, limit))
		}

		delta := 10.0 * 1000 * 1000 * 1000 //ns means 0.01s
		assert.InDelta(t, spend_time, time.Duration(timeout)*time.Second, delta, "[%s] Exectuion should take about %ds, but takes %ss (delta=%f)", test_name, timeout, spend_time.String(), delta)
	}
}

func TestPingSites(t *testing.T) {
	for_test_parametrized(t, 100, 3, context.Canceled)
	for_test_parametrized(t, 1000, 3, context.DeadlineExceeded)
	for_test_parametrized(t, 1000, 8, context.DeadlineExceeded)
}
