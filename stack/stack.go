package stack

import (
    "fmt"
)

func Init( )( stack *Stack_t ) {
    return &Stack_t{ }
}

func InitWithHandlers( ctx interface{ }, rlock_handler Rlock, runlock_handler Runlock, wlock_handler Wlock, unlock_handler Unlock )( stack *Stack_t ) {
    stack = Init( )

    stack.SetLockHandlers( ctx, rlock_handler, runlock_handler, wlock_handler, unlock_handler )

    return stack
}

func ( stack *Stack_t )SetLockHandlers( ctx interface{ }, rlock_handler Rlock, runlock_handler Runlock, wlock_handler Wlock, unlock_handler Unlock )( ) {
    if nil == stack {
        return
    }

    stack.Ctx             = ctx
    stack.Rlock_handler   = rlock_handler
    stack.Runlock_handler = runlock_handler
    stack.Wlock_handler   = wlock_handler
    stack.Unlock_handler  = unlock_handler
}

func ( stack *Stack_t )rlock( )( ) {
    if nil != stack && nil != stack.Rlock_handler {
        stack.Rlock_handler( stack.Ctx )
    }
}

func ( stack *Stack_t )runlock( )( ) {
    if nil != stack && nil != stack.Runlock_handler {
        stack.Runlock_handler( stack.Ctx )
    }
}

func ( stack *Stack_t )wlock( )( ) {
    if nil != stack && nil != stack.Wlock_handler {
        stack.Wlock_handler( stack.Ctx )
    }
}

func ( stack *Stack_t )unlock( )( ) {
    if nil != stack && nil != stack.Unlock_handler {
        stack.Unlock_handler( stack.Ctx )
    }
}

func ( stack *Stack_t )IsEmpty( )( bool ) {
    return nil != stack && len( stack.Values ) == 0
}

func ( stack *Stack_t )Push( value interface{ } )( nelems int, err error ) {
    if nil == stack {
        return 0, fmt.Errorf( "invalid stack %v", stack )
    }

    stack.wlock( )
    defer stack.unlock( )

    stack.Values = append( stack.Values, value )

    return len( stack.Values ), nil
}

func ( stack *Stack_t )Pop( )( value interface{ }, err error ) {
    if nil == stack {
        return 0, fmt.Errorf( "invalid stack %v", stack )
    }

    stack.wlock( )
    defer stack.unlock( )

    length := len( stack.Values )

    if length == 0 { 
        return nil, fmt.Errorf( "empty stack %v", stack )
    }

    value        = stack.Values[ length - 1 ]
    stack.Values = stack.Values[ : length - 1 ]

    return value, nil
}

func ( stack *Stack_t )Peek( )( value interface{ }, err error ) {
    if nil == stack {
        return 0, fmt.Errorf( "invalid stack %v", stack )
    }

    stack.rlock( )
    defer stack.runlock( )

    length := len( stack.Values )

    if length == 0 {
        return nil, fmt.Errorf( "empty stack %v", stack )
    }

    return stack.Values[ length - 1 ], nil
}
