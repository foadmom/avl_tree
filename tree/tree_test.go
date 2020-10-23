package tree

import (
	"fmt"
	"testing"
)

var keys_1	 []int    = []int {20,4,15,7,13,1,35};

var keys_2	 []int    = []int {20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1};

var keys_3	 []int    = []int {17,16,6,5,4,3,15,14,13,20,19,18,12,11,10,9,8,7,2,1};

var keys_4	 []int    = []int {1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,50,55,80,81,82,83,84,85,86,101,102,103,70,69,68,71,72,73,21,22,23,24,25,26};


func TestRandomTree (t *testing.T) {
	_slice := CreateRandomSlice (10, 0, 1000);
	fmt.Println (_slice);
}


func testCreateTestTree (nodeCount int, smallesKey int, largestKey int) *Tree {
	var _tree  Tree = Tree {nil, nil};
	var _value string = "node value";

	_keys := CreateRandomSlice (nodeCount, -smallesKey, largestKey);

	for _i := range _keys {
		var newNode *Node = CreateNode (_keys[_i], &_value);
		_tree.Insert (newNode);
	}
	return &_tree;
}


func TestTreeMeta (t *testing.T) {
	var  _tree    *Tree = testCreateTestTree (1000, -500, 20000);
	testNodeMeta (t, _tree.Root);
	testNodeKeys (t, _tree, keys_4);
}

func testNodeKeys (t *testing.T, tree *Tree, keys []int) {
	for _, _key := range keys {
		fmt.Printf ("looking for key %d\n", _key);
        if (tree.Find (_key) == nil) {
			t.Errorf ("key %d not found\n", _key);
		}
    }
}


func testNodeMeta (t *testing.T, node *Node) int {
	var _errorText string = CheckNodeMeta(node);
	if (_errorText != "") {
		t.Errorf("following errors were detected %s\n", _errorText);
	}
	return 0;
}



func TestNode_getBalance(t *testing.T) {
	type fields struct {
		Key     int
		Value   string
		Right   *Node
		Left    *Node
		Height  int
		Balance int
		Parent  *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parent := &Node{
				Key:     tt.fields.Key,
				Value:   tt.fields.Value,
				Right:   tt.fields.Right,
				Left:    tt.fields.Left,
				Height:  tt.fields.Height,
				Balance: tt.fields.Balance,
				Parent:  tt.fields.Parent,
			}
			if got := parent.getBalance(); got != tt.want {
				t.Errorf("Node.getBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}

