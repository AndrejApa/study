package main
import (
	"fmt"
	"sort"
)

func main() {

	m := map[int]string{2: "a", 0: "b", 1: "c"}

	values := make([]int, 0, len(m))
	for v:= range m {
		values = append(values, v)
	}
	sort.Ints(values)

	for _, v := range values {
		fmt.Println(v, m[v])
	}
}
