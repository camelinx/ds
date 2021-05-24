package stack

type Rlock func( interface{ } )( )
type Wlock func( interface{ } )( )
type Unlock func( interface{ } )( )

type Stack_t struct {
    Values          [ ]interface{ }

    Ctx                interface{ }
    Rlock_handler      Rlock
    Wlock_handler      Wlock
    Unlock_handler     Unlock
}
