package calcurator

// 小文字で始めた変数はプライベート変数となる(同一パッケージ内であれば参照可能)
var offset float64 = 1

// 大文字で始めた変数はグローバル変数となり外部から参照可能。
var Offset float64 = 1

func Sum(a float64, b float64) float64 {
	return a + b + offset
}
