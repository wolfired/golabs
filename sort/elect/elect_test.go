package elect

import "testing"

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
}
