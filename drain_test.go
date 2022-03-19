package lazy

import (
	"fmt"
	"go4ml.xyz/fu"
	"gotest.tools/v3/assert"
	"testing"
)

func Test_Drain_1(t *testing.T) {
	i := 0
	List(colors).MustDrain(Sink(func(v interface{}, _ error) (_ error) {
		if v != nil {
			x := v.(Color)
			assert.Assert(t, colors[i].Color == x.Color)
			i++
		} else {
			assert.Assert(t, i == len(colors))
		}
		return
	}))
}

func Test_Drain_2(t *testing.T) {
	i := 0
	List(colors).MustDrain(Sink(func(v interface{}, _ error) (_ error) {
		if v != nil {
			x := v.(Color)
			assert.Assert(t, colors[i].Color == x.Color)
			i++
		} else {
			assert.Assert(t, i == len(colors))
		}
		return
	}), 8)
}

func Test_CcrDrain_1(t *testing.T) {
	c := fu.AtomicCounter{}
	List(colors).MustDrain(func(_ int) []Worker {
		wrk := make([]Worker, 8) // concurrency = 8
		for i := range wrk {
			wrk[i] = func(i int, v interface{}, _ error) error {
				if v != nil {
					switch x := v.(type) {
					case Color:
						assert.Assert(t, colors[i].Color == x.Color)
						i++
						c.Inc()

					}
				}
				return nil
			}
		}
		return wrk
	})
	fmt.Println(c.Value)
	assert.Assert(t, int(c.Value) == len(colors))
}
