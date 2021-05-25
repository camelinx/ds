package bst

type BstNodeCallback func( interface{ }, interface{ } )( error )
type Comparator func( interface{ }, interface{ } )( int8, error )

type Rlock func( interface{ } )( )
type Runlock func( interface{ } )( )
type Wlock func( interface{ } )( )
type Unlock func( interface{ } )( )

type BstNode_t struct {
    Value            interface{ }

    Parent          *BstNode_t
    Left            *BstNode_t
    Right           *BstNode_t
}

type Bst_t struct {
    Root            *BstNode_t
    Count            uint

    Ctx              interface{ }
    Rlock_handler    Rlock
    Runlock_handler  Runlock
    Wlock_handler    Wlock
    Unlock_handler   Unlock
}
