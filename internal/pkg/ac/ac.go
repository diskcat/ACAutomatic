package ac

import "math"

type AC struct {
	children map[rune]*AC
	isEnd    bool
	fail     *AC
}

func NewAC(content []string) *AC {
	ac := &AC{children: map[rune]*AC{}}
	creatTrie(ac, content)
	buildFail(ac)
	return ac
}

func creatTrie(ac *AC, content []string) {
	for _, sentence := range content {
		words := []rune(sentence)
		iter := ac
		for _, word := range words {
			if _, ok := iter.children[word]; !ok {
				iter.children[word] = &AC{children: map[rune]*AC{}}
			}
			iter = iter.children[word]
		}
		iter.isEnd = true
	}
}

func buildFail(ac *AC) {
	queue := []*AC{ac}
	for len(queue) > 0 {
		parent := queue[0]
		queue = queue[1:]
		for word, child := range parent.children {
			if parent == ac {
				child.fail = ac
			} else {
				if parent.fail != nil {
					if _, ok := parent.fail.children[word]; ok {
						child.fail = parent.fail.children[word]
					} else {
						child.fail = parent.fail
					}
				} else {
					child.fail = ac
				}
			}
			queue = append(queue, child)
		}
	}
}

func (ac *AC) Insert(sentence string) {
	node := ac
	for _, word := range sentence {
		if _, ok := node.children[word]; !ok {
			if node == ac {
				node.children[word] = &AC{children: map[rune]*AC{}, fail: ac}
			} else {
				node.children[word] = &AC{children: map[rune]*AC{}, fail: node.fail}
				if _, ok := node.fail.children[word]; ok {
					node.children[word].fail = node.fail.children[word]
				}
			}
		}
		node = node.children[word]
	}
	node.isEnd = true
}

//status: 0删除成功，1删除失败
func (ac *AC) Delete(line string) (isDelete bool) {
	_line := []rune(line)
	up := math.MinInt32
	n := len(_line)
	hit := false
	var dfs func(ac *AC, idx int)
	dfs = func(ac *AC, idx int) {
		if idx == n {
			return
		}
		if _, ok := ac.children[_line[idx]]; !ok {
			return
		}
		dfs(ac.children[_line[idx]], idx+1)
		if ac.children[_line[idx]].isEnd && idx < n-1 && !hit {
			up = idx
			hit = true
		}
		if idx <= up {
			return
		}
		if len(ac.children[_line[idx]].children) > 1 && !hit {
			hit = true
			return
		}
		if !hit && len(ac.children[_line[idx]].children) == 0 {
			delete(ac.children, _line[idx])
			isDelete = true
		}
	}
	dfs(ac, 0)
	return
}

func (ac *AC) Search(content string) (ans []string) {
	words := []rune(content)
	iter := ac
	var begin, end int
	for i, word := range words {
		_, ok := iter.children[word]
		for !ok && iter != ac {
			iter = iter.fail
		}
		_, ok = iter.children[word]
		if ok {
			if iter == ac {
				begin = i
			}
			iter = iter.children[word]
			if iter.isEnd {
				end = i
				ans = append(ans, string(words[begin:end+1]))
			}
		}
	}
	return
}

func (ac *AC) Replace(content string) string {
	words := []rune(content)
	iter := ac
	var begin, end int
	for i, word := range words {
		_, ok := iter.children[word]
		for !ok && iter != ac {
			iter = iter.fail
		}
		_, ok = iter.children[word]
		if ok {
			if iter == ac {
				begin = i
			}
			iter = iter.children[word]
			if iter.isEnd {
				end = i
				for j := begin; j <= end; j++ {
					words[j] = '*'
				}
				begin = end
			}
		}
	}
	return string(words)
}
