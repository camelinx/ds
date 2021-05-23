package bst

import (
    "fmt"
)

func Init( )( tree *Bst_t ) {
    return &Bst_t{ Root : nil, Count : 0 }
}

func ( tree *Bst_t )Insert( value interface{ }, comparator Comparator )( count uint, err error ) {
    if nil == tree {
        return 0, fmt.Errorf( "failed to insert into nil tree" )
    }

    new := &BstNode_t{ Value : value }

    if nil == tree.Root {
        tree.Root  = new
        tree.Count = 1
    } else {
        cur := tree.Root
        for nil != cur {
            cmp, err := comparator( value, cur.Value )
            if nil != err {
                return 0, fmt.Errorf( "failed to insert into tree, comparison failed with error %v", err )
            }

            new.Parent = cur

            if 0 == cmp {
                break
            } else if cmp < 0 {
                if nil == cur.Left {
                    cur.Left = new
                    tree.Count++

                    break
                }

                cur = cur.Left
            } else {
                if nil == cur.Right {
                    cur.Right = new
                    tree.Count++

                    break
                }

                cur = cur.Right
            }
        }
    }

    return tree.Count, nil
}

func ( tree *Bst_t )findNode( value interface{ }, comparator Comparator )( node *BstNode_t, err error ) {
    if nil == tree || nil == tree.Root {
        return nil, fmt.Errorf( "failed to insert into nil tree" )
    }

    cur := tree.Root
    for nil != cur {
        cmp, err := comparator( value, cur.Value )
        if nil != err {
            return nil, fmt.Errorf( "failed to search in tree, comparison failed with error %v", err )
        }

        if 0 == cmp {
            return cur, nil
        } else if cmp < 0 {
            cur = cur.Left
        } else {
            cur = cur.Right
        }
    }

    return nil, nil
}

// For leaf nodes, parent is the neighbor
// For intermediate nodes, if right sub tree exists, left most child of right sub tree else the right most child of left sub tree
func ( tree *Bst_t )findNeighbor( node *BstNode_t, comparator Comparator )( neighbor *BstNode_t, err error ) {
    if nil == tree || nil == node {
        return nil, fmt.Errorf( "failed to find neighbor, either tree or node is nil" )
    }

    if nil == node.Left && nil == node.Right {
        return node.Parent, nil
    }

    if nil != node.Right {
        neighbor = node.Right
        for nil != neighbor.Left {
            neighbor = neighbor.Left
        }
    } else {
        neighbor = node.Left
        for nil != neighbor.Right {
            neighbor = neighbor.Right
        }
    }

    return neighbor, nil
}

func ( tree *Bst_t )Search( value interface{ }, comparator Comparator )( found bool, err error ) {
    node, err := tree.findNode( value, comparator )
    if nil == err && nil != node {
        return true, nil
    }

    return false, err
}

func ( tree *Bst_t )removeLeafNode( node *BstNode_t )( err error ) {
    if nil != node.Left || nil != node.Right {
        return fmt.Errorf( "%v not a leaf node", node )
    }

    if nil == node.Parent {
        return fmt.Errorf( "%v has no parent reference", node )
    }

    if node.Parent.Left == node {
        node.Parent.Left = nil
        tree.Count--
    } else if node.Parent.Right == node {
        node.Parent.Right = nil
        tree.Count--
    } else {
        return fmt.Errorf( "%v has an invalid parent reference", node )
    }

    return nil
}

func ( tree *Bst_t )Delete( value interface{ }, comparator Comparator )( count uint, err error ) {
    node, err := tree.findNode( value, comparator )
    if nil != err || nil == node {
        return tree.Count, err
    }

    neighbor, err := tree.findNeighbor( node, comparator )
    if nil != err {
        return tree.Count, err
    }

    for {
        if nil == node.Left && nil == node.Right {
            err = tree.removeLeafNode( node )
            if nil != err {
                return tree.Count, err
            }

            break
        }

        node.Value = neighbor.Value
        node       = neighbor

        neighbor, err = tree.findNeighbor( node, comparator )
        if nil != err {
            return tree.Count, err
        }
    }

    return tree.Count, nil
}
