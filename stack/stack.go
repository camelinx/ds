package stack

import (
    "fmt"
)

func Init( )( stack *Stack_t ) {
    return &Stack_t{ }
}

func ( stack *Stack_t )push( value interface{ } )( nelems int, err error ) {
    if nil == stack {
        return 0, fmt.Errorf( "invalid stack %v", stack )
    }

    stack.Values = append( stack.Values, value )

    return len( stack.Values ), nil
}

func ( stack *Stack_t )pop( )( value interface{ }, err error ) {
    if nil == stack  || 0 == len( stack.Values ) {
        return 0, fmt.Errorf( "invalid stack %v", stack )
    }

    value = stack.Values[ len( stack.Values ) - 1 ]
    stack.Values = stack.Values[ : len( stack.Values ) - 1 ]

    return value, nil
}

func ( stack *Stack_t )peek( )( value interface{ }, err error ) {
    if nil == stack || 0 == len( stack.Values ) {
        return 0, fmt.Errorf( "invalid stack %v", stack )
    }

    return stack.Values[ len( stack.Values ) - 1 ], nil
}
