package main

import "fmt"

// Trie ...
type Trie struct {
	r byte
	e bool
	m map[rune]*Trie
}

// Constructor ...
func Constructor() Trie {
	return Trie{m: make(map[rune]*Trie)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, v := range word {
		t, ok := this.m[v]
		if !ok {
			t = &Trie{m: make(map[rune]*Trie)}
			this.m[v] = t
		}
		this = t
	}
	this.e = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, v := range word {
		t, ok := this.m[v]
		if !ok {
			return false
		}
		this = t
	}
	return this.e
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		t, ok := this.m[v]
		if !ok {
			return false
		}
		this = t
	}
	return true
}

func main() {
	t := Constructor()
	t.Insert("apple")
	fmt.Println(t.Search("apple"))
	fmt.Println(t.StartsWith("apple"))
	fmt.Println(t.StartsWith("apples"))
}
