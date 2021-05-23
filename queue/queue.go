package queue

import (
    "fmt"
)

func Init( )( queue *Queue_t ) {
    return &Queue_t{ }
}

func ( queue *Queue_t )push( value interface{ } )( nelems int, err error ) {
    if nil == queue {
        return 0, fmt.Errorf( "invalid queue %v", queue )
    }

    queue.Values = append( queue.Values, value )

    return len( queue.Values ), nil
}

func ( queue *Queue_t )pop( )( value interface{ }, err error ) {
    if nil == queue || 0 == len( queue.Values ) {
        return 0, fmt.Errorf( "invalid queue %v", queue )
    }

    value = queue.Values[ 0 ]
    queue.Values = queue.Values[ 1 : len( queue.Values ) ]

    return value, nil
}

func ( queue *Queue_t )peek( )( value interface{ }, err error ) {
    if nil == queue || 0 == len( queue.Values ) {
        return 0, fmt.Errorf( "invalid queue %v", queue )
    }

    return queue.Values[ 0 ], nil
}
