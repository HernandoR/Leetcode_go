package main

/*
   Author  :   Liuzhen
   Ques of     [2039]the-time-when-the-network-becomes-idle
   Create  :   2022-03-20 11:02:44	获取当前时间
*/

func main() {

	//ans := networkBecomesIdle()
	//fmt.Println(ans)
}

//leetcode submit region begin(Prohibit modification and deletion)
func networkBecomesIdle(edges [][]int, patience []int) int {

	return 0
}

func (g unDirectGraph) AddEdge(Node1, Node2 node) {

}

func (g unDirectGraph) AddNode(id int, patience int) {

	g.Nodes[id] = node{
		ID:          id,
		Patience:    patience,
		LinkedNodes: nil,
	}
}

type unDirectGraph struct {
	//Nodes []node
	Nodes map[int]node
	Edges map[string]edge
}

type node struct {
	ID          int
	Patience    int
	LinkedNodes []*node
}

type edge struct {
	NodeHoster node
	NodeClient node
}

//leetcode submit region end(Prohibit modification and deletion)
