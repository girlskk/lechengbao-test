package main

import "fmt"

// Perm() 对 a 形成的每⼀排列调⽤ f().
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// 对索引 i 从 0 到 len(a) - 1，实现递归函数 perm().
func perm(a []rune, f func([]rune), i int) {
	if i == len(a)-1 {
		f(a)
	}

	for s := i; s < len(a); s++ {
		swap(a, i, s)
		perm(a, f, i+1)
		swap(a, i, s)
	}
}

func swap(a []rune, i, j int) {
	if i > len(a)-1 || j > len(a)-1 {
		return
	}
	a[i], a[j] = a[j], a[i]
}
func main() {
	Perm([]rune("ABC"), func(a []rune) {
		fmt.Println(string(a))
	})
}
