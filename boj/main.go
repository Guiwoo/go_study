package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(len(" "))
}

func solution(merchantNames []string) []string {
	type real struct {
		name    string
		space   int
		special bool
	}

	spaceCount := func(s string) (int, bool) {
		special := []rune{'&', '(', ')', '.', ',', '-'}
		cnt := 0
		spec := false
		for _, v := range s {
			if v == ' ' {
				cnt++
			}
			for _, r := range special {
				if v == r {
					spec = true
					break
				}
			}
		}
		return cnt, spec
	}

	var answer []string
	reals := make([]*real, len(merchantNames))
	for i, v := range merchantNames {
		space, spec := spaceCount(v)
		reals[i] = &real{
			v,
			len(v) - space,
			spec,
		}
	}

	sort.Slice(reals, func(i, j int) bool {
		if len(reals[i].name)-reals[i].space != len(reals[j].name)-reals[j].space {
			return len(reals[i].name)-reals[i].space > len(reals[j].name)-reals[j].space
		}
		if reals[i].special && !reals[j].special {
			return true
		} else if reals[j].special && !reals[i].special {
			return false
		}
		return true
	})

	for _, v := range reals {
		fmt.Println(v)
	}
	fmt.Println(answer)
	return nil
}
