package bst

type BstNodeIterator func( interface{ }, interface{ } )( error )
type Comparator func( interface{ }, interface{ } )( int8, error )

type BstNode_t struct {
    Value       interface{ }

    Parent     *BstNode_t
    Left       *BstNode_t
    Right      *BstNode_t
}

type Bst_t struct {
    Root       *BstNode_t

    Count       uint
}
