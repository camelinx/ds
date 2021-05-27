package sort

import (
    "fmt"
    "reflect"

    "github.com/ds/stack"
)

type part_t struct {
    start   int
    end     int
}

func convertArgToSlice( arg interface{ } )( val reflect.Value, ok bool ) {
    val = reflect.ValueOf( arg )
    if reflect.Slice == val.Kind( ) {
        ok = true
    }

    return val, ok
}

func Qsort( arg interface{ }, compare Comparator )( result [ ]interface{ }, err error ) {
    if nil == compare || nil == arg {
        return nil, fmt.Errorf( "invalid arguments" )
    }

    argSlice, ok := convertArgToSlice( arg )
    if !ok {
        return nil, fmt.Errorf( "argument is not a slice" )
    }

    for i := 0; i < argSlice.Len( ); i++ {
        result = append( result, argSlice.Index( i ).Interface( ) )
    }

    if argSlice.Len( ) < 2 {
        return result, nil
    }

    parts := stack.Init( )

    start := 0
    end   := len( result ) - 1

    parts.Push( part_t{ start : start, end : end } )

    for !parts.IsEmpty( ) {
        pop, err := parts.Pop( )
        if err != nil {
            return nil, fmt.Errorf( "corrupted stack" )
        }

        part  := pop.( part_t )
        mid   := ( part.start + part.end ) / 2

        left  := part.start
        right := part.end

        if 2 > ( right - left ) {
            cmp, err := compare( result[ left ], result[ right ] )
            if err != nil {
                return nil, fmt.Errorf( "comparison failed with error %v", err )
            }

            if cmp > 0 {
                result[ left ], result[ right ] = result[ right ], result[ left ]
            }

            continue
        }

        result[ right ], result[ mid ] = result[ mid ], result[ right ]

        for i := left; i < right; i++ {
            cmp, err := compare( result[ i ], result[ right ] )
            if err != nil {
                return nil, fmt.Errorf( "comparison failed with error %v", err )
            }

            if cmp < 0 {
                result[ left ], result[ i ] = result[ i ], result[ left ]
                left++
            }
        }

        result[ left ], result[ right ] = result[ right ], result[ left ]

        if part.start < left {
            parts.Push( part_t{ start : part.start, end : left } )
        }

        if left + 1 < part.end {
            parts.Push( part_t{ start : left + 1, end : part.end } )
        }
    }

    return result, nil
}
