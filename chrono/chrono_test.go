package chrono

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeNow(t *testing.T) {
	t.Parallel()
	c := NewTimeNow()
	c.Update()

	now := time.Now().UTC()
	cnow := time.Unix(c.Get(), 0)
	assert.WithinDuration(t, now, cnow, 2*time.Second)
}
