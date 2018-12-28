package shell

import (
	"math/rand"
	"testing"
	"time"
)

func Test_optimize(t *testing.T) {
	{
		sample := [...]int{}
		optimize(sample[:])
		want := [...]int{}
		if want != sample {
			t.Error("Want", want)
			t.Error("But", sample)
		}
	}

	{
		sample := [...]int{0}
		optimize(sample[:])
		want := [...]int{0}
		if want != sample {
			t.Error("Want", want)
			t.Error("But", sample)
		}
	}

	{
		sample := [...]int{0, 1}
		optimize(sample[:])
		want := [...]int{0, 1}
		if want != sample {
			t.Error("Want", want)
			t.Error("But", sample)
		}
	}

	{
		sample := [...]int{1, 0}
		optimize(sample[:])
		want := [...]int{0, 1}
		if want != sample {
			t.Error("Want", want)
			t.Error("But", sample)
		}
	}

	{
		sample := [...]int{0, 0, 1, 1}
		optimize(sample[:])
		want := [...]int{0, 0, 1, 1}
		if want != sample {
			t.Error("Want", want)
			t.Error("But", sample)
		}
	}

	{
		sample := [...]int{1, 1, 0, 0}
		optimize(sample[:])
		want := [...]int{0, 0, 1, 1}
		if want != sample {
			t.Error("Want", want)
			t.Error("But", sample)
		}
	}

	{
		sample := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
		optimize(sample[:])
		want := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		if want != sample {
			t.Error("Want", want)
			t.Error("But", sample)
		}
	}

	rand.Seed(time.Now().UnixNano())
	sample := rand.Perm(100000)
	optimize(sample[:])

	for j := 1; j < len(sample); j++ {
		if sample[j-1] > sample[j] {
			t.Error("Fault", sample[j-1], sample[j])
			break
		}
	}
}

func Benchmark_optimize(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		rand.Seed(time.Now().UnixNano())
		sample := rand.Perm(100000)
		b.StartTimer()
		optimize(sample)
	}
}
