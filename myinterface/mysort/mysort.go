package mysort

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (p StringSlice) Len() int { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func Main() {
  //fmt.Println(sort.Sort(StringSlice("Yusuf")))
  name := []string{"e", "d", "c", "b", "a"}
  //sort.Strings(name)
  sort.Sort(StringSlice(name))
  fmt.Println("name", name)
}