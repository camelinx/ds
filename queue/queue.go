package queue

import (
    "fmt"
)

func Init( )( queue *Queue_t ) {
    return &Queue_t{ }
}

func InitWithHandlers( ctx interface{ }, rlock_handler Rlock, runlock_handler Runlock, wlock_handler Wlock, unlock_handler Unlock )( queue *Queue_t ) {
    queue = Init( )

    queue.SetLockHandlers( ctx, rlock_handler, runlock_handler, wlock_handler, unlock_handler )

    return queue
}

func ( queue *Queue_t )SetLockHandlers( ctx interface{ }, rlock_handler Rlock, runlock_handler Runlock, wlock_handler Wlock, unlock_handler Unlock )( ) {
    if nil == queue {
        return
    }

    queue.Ctx              = ctx
    queue.Rlock_handler    = rlock_handler
    queue.Runlock_handler  = runlock_handler
    queue.Wlock_handler    = wlock_handler
    queue.Unlock_handler   = unlock_handler
}

func ( queue *Queue_t )rlock( )( ) {
    if nil != queue && nil != queue.Rlock_handler {
        queue.Rlock_handler( queue.Ctx )
    }
}

func ( queue *Queue_t )runlock( )( ) {
    if nil != queue && nil != queue.Runlock_handler {
        queue.Runlock_handler( queue.Ctx )
    }
}

func ( queue *Queue_t )wlock( )( ) {
    if nil != queue && nil != queue.Wlock_handler {
        queue.Wlock_handler( queue.Ctx )
    }
}

func ( queue *Queue_t )unlock( )( ) {
    if nil != queue && nil != queue.Unlock_handler {
        queue.Unlock_handler( queue.Ctx )
    }
}

func ( queue *Queue_t )IsEmpty( )( bool ) {
    return nil != queue && len( queue.Values ) == 0
}

func ( queue *Queue_t )Push( value interface{ } )( nelems int, err error ) {
    if nil == queue {
        return 0, fmt.Errorf( "invalid queue %v", queue )
    }

    queue.wlock( )
    defer queue.unlock( )

    queue.Values = append( queue.Values, value )

    return len( queue.Values ), nil
}

func ( queue *Queue_t )Pop( )( value interface{ }, err error ) {
    if nil == queue {
        return 0, fmt.Errorf( "invalid queue %v", queue )
    }

    queue.wlock( )
    defer queue.unlock( )

    length := len( queue.Values )

    if length == 0 { 
        return nil, fmt.Errorf( "empty queue %v", queue )
    }

    value        = queue.Values[ 0 ]
    queue.Values = queue.Values[ 1 : length ]

    return value, nil
}

func ( queue *Queue_t )Peek( )( value interface{ }, err error ) {
    if nil == queue {
        return 0, fmt.Errorf( "invalid queue %v", queue )
    }

    queue.rlock( )
    defer queue.runlock( )

    if len( queue.Values ) == 0 {
        return nil, fmt.Errorf( "empty queue %v", queue )
    }

    return queue.Values[ 0 ], nil
}
