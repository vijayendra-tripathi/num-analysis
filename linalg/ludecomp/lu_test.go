package ludecomp

import (
	"math/rand"
	"testing"

	"github.com/unixpickle/num-analysis/linalg"
)

func TestSolve4x4(t *testing.T) {
	m := &linalg.Matrix{
		Rows: 4,
		Cols: 4,
		Data: []float64{
			1, 2, 3, 7,
			4, 5, 6, 9.5,
			7, 8, 10, 3.2,
			1.4, 1.5, 7.9, 2.1,
		},
	}
	dec := Decompose(m)

	inputs := []linalg.Vector{
		{1, 2, 3, 7},
		{1, 2, 5, 2},
	}
	expected := []linalg.Vector{
		{2075.0 / 1589.0, -512.0 / 245.0, 7769.0 / 7612.0, 397.0 / 3423.0},
		{-223.0 / 193.0, 1329.0 / 941.0, 1925.0 / 7927.0, -371.0 / 1858.0},
	}
	for i, in := range inputs {
		out := dec.Solve(in)
		if vectorDiff(out, expected[i]) > 0.0001 {
			t.Error("Test", i, "expected", expected[i], "but got", out)
		}
	}
}

func BenchmarkDecompose200x200(b *testing.B) {
	mat := randMatrix(200)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Decompose(mat)
	}
}

func BenchmarkDecompose100x100(b *testing.B) {
	mat := randMatrix(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Decompose(mat)
	}
}

func BenchmarkSolve800x800(b *testing.B) {
	lu := Decompose(randMatrix(800))
	vec := randVec(800)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lu.Solve(vec)
	}
}

func BenchmarkSolve400x400(b *testing.B) {
	lu := Decompose(randMatrix(400))
	vec := randVec(400)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lu.Solve(vec)
	}
}

func randMatrix(size int) *linalg.Matrix {
	res := linalg.NewMatrix(size, size)
	for i := 0; i < res.Rows; i++ {
		for j := 0; j < res.Rows; j++ {
			res.Set(i, j, rand.Float64())
		}
	}
	return res
}

func randVec(size int) linalg.Vector {
	res := make(linalg.Vector, size)
	for i := 0; i < size; i++ {
		res[i] = rand.Float64()
	}
	return res
}
