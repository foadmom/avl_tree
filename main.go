package main

import (
	t "avl_tree/tree"
	"fmt"
)

var tree  t.Tree;

func main () {
	fmt.Println ("btree starting");

	tempTest(&tree);
}


func tempTest (tree *t.Tree) {
	_keys := t.CreateRandomSlice (1000, -50000, 50000);
//	_keys := []int {5,4,6,3,7,2,8,1,9};
//	var _keys	 []int    = []int {1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,50,55,80,81,82,83,84,85,86,101,102,103,70,69,68,71,72,73,21,22,23,24,25,26};

	fmt.Printf ("size of the input nodes = %d\n", len(_keys));
	for _i := range _keys {
		var _value  string = fmt.Sprintf ("value-%d", _keys[_i]);

		var newNode *t.Node = t.CreateNode (_keys[_i], &_value);
		tree.Insert (newNode);
//		tree.Root.Dump(0, "");
		fmt.Printf (t.CheckNodeMeta (tree.Root));
	}

	var _count int;
	tree.Root.CountNodes (&_count);
	fmt.Printf ("total node count = %d\n", _count);
	var  _treeHeight     int = tree.Root.CalculateHeight();
	fmt.Printf ("tree height = %d\n", _treeHeight);
	fmt.Printf ("smallest value = %d\n", tree.Root.FindSmallestOnLeft().Key);
	fmt.Printf ("larest   value = %d\n", tree.Root.FindLargestOnRight().Key);
	tree.Delete (69);
	tree.Delete (26);
//	tree.Root.Dump(0, "");
	tree.Update (103, "******************")
//	tree.Root.Dump(0, "");
}



