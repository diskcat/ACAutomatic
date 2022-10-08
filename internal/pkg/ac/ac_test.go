package ac

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestNewAC(t *testing.T) {
	content := []string{"he", "she", "his", "her"}
	ac := NewAC(content)
	formatPrintAC(ac)
}

func formatPrintAC(ac *AC) {
	fmt.Println(unsafe.Pointer(ac))
	fails := [][]*AC{}
	fail := []*AC{}
	ans := [][]*AC{}
	path := []*AC{}
	strs := []string{}
	alphs := []rune{}
	var dfs func(ac *AC)
	dfs = func(ac *AC) {
		if ac.isEnd && len(ac.children) == 0 {
			ans = append(ans, append([]*AC{}, path...))
			fails = append(fails, append([]*AC{}, fail...))
			strs = append(strs, string(alphs))
			return
		}
		for k, v := range ac.children {
			path = append(path, v)
			fail = append(fail, v.fail)
			alphs = append(alphs, k)
			dfs(v)
			alphs = alphs[:len(alphs)-1]
			fail = fail[:len(fail)-1]
			path = path[:len(path)-1]
		}
	}
	dfs(ac)
	fmt.Println(ans)
	fmt.Println(fails)
	fmt.Println(strs)
}

func TestInsert(t *testing.T) {
	content := []string{"he", "she", "his", "her"}
	ac := &AC{children: map[rune]*AC{}}
	for _, sentence := range content {
		ac.Insert(sentence)
	}
	formatPrintAC(ac)
}

func TestReplace(t *testing.T) {
	pattern := "he and she are good friends! hers"
	content := []string{"he", "she", "his", "her"}
	ac := NewAC(content)
	fmt.Println(ac.Replace(pattern))
}

func TestSearch(t *testing.T) {
	pattern := "he and she are good friends! hers"
	content := []string{"he", "she", "his", "her"}
	ac := NewAC(content)
	fmt.Println(ac.Search(pattern))
}

func TestDelete(t *testing.T) {
	pattern := "he and she are good friends! hers"
	content := []string{"he", "she", "his", "her"}
	ac := NewAC(content)
	ac.Delete("her")
	fmt.Println(ac.Search(pattern))
}
