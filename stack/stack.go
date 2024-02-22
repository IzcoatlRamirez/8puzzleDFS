package stack

import (
    "github.com/IzcoatlRam/puzzle8dfs/node"
    "container/list"
)

type Stack struct {
    s *list.List
}

func NewStack() *Stack {
    return &Stack{
        s: list.New(),
    }
}

func (s *Stack) Push(n *node.Node) {
    s.s.PushFront(n)
}

func (s *Stack) Pop() *node.Node {
    e := s.s.Front()
    s.s.Remove(e)
    return e.Value.(*node.Node)
}

func (s *Stack) Len() int {
    return s.s.Len()
}