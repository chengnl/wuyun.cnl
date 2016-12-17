package consistenthash

import (
	"errors"
	"fmt"

	"github.com/emirpasic/gods/trees/redblacktree"
)

//FNV1A_32_HASH
const (
	FNV_32_INIT  int = 2166136261
	FNV_32_PRIME int = 16777619
)

type ConsistenHash struct {
	vNum    int
	members *redblacktree.Tree
	hfunc   func(string) int
}

func NewConsistenHash(virtualNum int, nodes []interface{}) *ConsistenHash {
	chash := &ConsistenHash{vNum: virtualNum, members: redblacktree.NewWithIntComparator(), hfunc: func(hashString string) int {
		hashCode := FNV_32_INIT
		for _, r := range hashString {
			hashCode = (hashCode ^ int(r)) * FNV_32_PRIME
		}
		//fmt.Printf("hashCode:%d\n", hashCode&0xffffffff)
		return hashCode & 0xffffffff
	}}
	for _, node := range nodes {
		chash.AddNode(node)
	}
	return chash
}
func NewConsistenHashWithHFunc(virtualNum int, hashfunc func(string) int, nodes []interface{}) *ConsistenHash {
	chash := &ConsistenHash{vNum: virtualNum, members: redblacktree.NewWithIntComparator(), hfunc: hashfunc}
	for _, node := range nodes {
		chash.AddNode(node)
	}
	return chash
}
func (chash *ConsistenHash) AddNode(node interface{}) {
	nodeString, err := chash.getNodeString(node)
	if err != nil {
		return
	}
	for i := 0; i < chash.vNum; i++ {
		nodeString = fmt.Sprintf("%s-%d", nodeString, i)
		chash.members.Put(chash.hfunc(nodeString), node)
	}
}
func (chash *ConsistenHash) DeleteNode(node interface{}) {
	nodeString, err := chash.getNodeString(node)
	if err != nil {
		return
	}
	for i := 0; i < chash.vNum; i++ {
		nodeString = fmt.Sprintf("%s-%d", nodeString, i)
		chash.members.Remove(chash.hfunc(nodeString))
	}
}
func (chash *ConsistenHash) GetNode(objStr string) interface{} {
	key := chash.hfunc(objStr)
	node, found := chash.members.Ceiling(key)
	if found {
		return node.Value
	}
	if chash.members.Empty() {
		return nil
	}
	return chash.members.Values()[0]
}

func (chash *ConsistenHash) getNodeString(node interface{}) (string, error) {
	switch node.(type) {
	case nil:
		return "", errors.New("delete node is nil")
	case int, uint:
		return fmt.Sprintf("%d", node), nil
	case string:
		return node.(string), nil
	case Node:
		return node.(Node).String(), nil
	default:
		return fmt.Sprintf("%v", node), nil
	}
}
