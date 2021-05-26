package sort

import (
    "fmt"

    "github.com/ds/stack"
)

type part_t struct {
    start   int
    end     int
}

func qsort( values [ ]interface{ }, compare Comparator )( err error ) {
    if nil == compare {
        return fmt.Errorf( "invalid comparator" )
    }

    if nil == values || len( values ) < 2 {
        return nil
    }

    parts := stack.Init( )

    start := 0
    end   := len( values ) - 1
    
    parts.Push( part_t{ start : start, end : end } )

    for !parts.IsEmpty( ) {
        pop, err := parts.Pop( )
        if err != nil {
            return fmt.Errorf( "corrupted stack" )
        }

        part  := pop.( part_t )
        mid   := ( part.start + part.end ) / 2

        left  := part.start
        right := part.end

        values[ right ], values[ mid ] = values[ mid ], values[ right ]

        for i := 0; i < right; i++ {
            cmp, err := compare( values[ i ], values[ right ] )
            if err != nil {
                return fmt.Errorf( "comparison failed with error %v", err )
            }

            if cmp < 0 {
                values[ left ], values[ i ] = values[ i ], values[ left ]
                left++
            }
        }

        parts.Push( part_t{ start : part.start, end : left - 1 } )
        parts.Push( part_t{ start : left, end : part.end } )
    }

    return nil
}
