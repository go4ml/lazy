package lazy

import (
	"gotest.tools/v3/assert"
	"reflect"
	"sudachen.xyz/pkg/errstr"
	"testing"
)

func linlist(list interface{}) Source {
	return func(xs ...interface{}) Stream {
		worker := 0
		open := NoPrefetch
		for _, x := range xs {
			if f, ok := x.(func()(int,int,Prefetch)); ok {
				worker, _, open = f()
			} else {
				return Error(errstr.Errorf("unsupported source option: %v", x))
			}
		}
		v := reflect.ValueOf(list)
		return open(worker,func()Stream{
			index := 0
			return func(next bool) (r interface{}, i int) {
				if next && index < v.Len() {
					i = index
					r = v.Index(i).Interface()
					index++
					return r, i
				}
				return EoS, index
			}
		})
	}
}

func Test_Prefetch_1(t *testing.T) {
	i := 0
	linlist(colors).MustDrain(Sink(func(v interface{}, _ error)(_ error){
		if v != nil {
			x := v.(Color)
			y := colors[i]
			assert.Assert(t, y.Color == x.Color)
			i++
		} else {
			assert.Assert(t, i == len(colors))
		}
		return
	}),2)
}

func Test_Prefetch_2(t *testing.T) {
	a := Sequence(func(i int)interface{}{
		if i < 100 { return i }
		return EoS
	}).MustCollectAny(2).([]int)
	assert.Assert(t, len(a) == 100)
}
