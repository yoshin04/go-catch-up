package calculator

func Multiply(a float64, b float64) float64 {
	return a * b
}

func multiply(a float64, b float64) float64 {
	// 同じパッケージ内であれば、小文字で始まる関数も呼び出せる変数も呼び出せる
	// ただし、他のパッケージからは呼び出せない
	return (a * b) + offset
}
