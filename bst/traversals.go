package bst

import (
    "fmt"

    stack_impl "github.com/ds/stack"
)

func ( tree *Bst_t )InOrderTraversal( callback BstNodeCallback )( err error ) {
    if nil == tree || nil == tree.Root {
        return fmt.Errorf( "invalid tree" )
    }

    stack := stack_impl.InitWithHandlers(
        tree.Ctx,
        stack_impl.Rlock( tree.Rlock_handler ),
        stack_impl.Runlock( tree.Runlock_handler ),
        stack_impl.Wlock( tree.Wlock_handler ),
        stack_impl.Unlock( tree.Unlock_handler ),
    )

    if nil == stack {
        return fmt.Errorf( "failed to initialize" )
    }

    stack.Push( tree.Root )

    var cur_node *BstNode_t
    var ok bool

    for !stack.IsEmpty( ) {
        if peek, err := stack.Peek( ); err != nil {
            cur_node, ok = peek.( *BstNode_t )
            if !ok {
                return fmt.Errorf( "corrupted tree" )
            }
        }

        if nil != cur_node.Left {
            stack.Push( cur_node.Left )
            continue
        }

        for {
            if nil != callback {
                err := callback( tree.Ctx, cur_node.Value )
                if nil != err {
                    return fmt.Errorf( "terminating: callback error = %v", err )
                }
            }

            stack.Pop( )

            if nil != cur_node.Right {
                stack.Push( cur_node.Right )
                break
            }

            if stack.IsEmpty( ) {
                break
            }

            if peek, err := stack.Peek( ); err != nil {
                cur_node, ok = peek.( *BstNode_t )
                if !ok {
                    return fmt.Errorf( "corrupted tree" )
                }
            }
        }
    }

    return nil
}
