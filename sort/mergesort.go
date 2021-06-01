package sort

import (
    "fmt"
)

func mergeParts( arg [ ]interface{ }, compare Comparator, part1 part_t, part2 part_t )( err error ) {
    var result [ ]interface{ }

    i := part1.start
    if i > part2.start {
        i = part2.start
    }

    for part1.start <= part1.end || part2.start <= part2.end {
        if part1.start > part1.end {
            for part2.start <= part2.end {
                result = append( result, arg[ part2.start ] )
                part2.start++
            }

            break
        }

        if part2.start > part2.end {
            for part1.start <= part1.end {
                result = append( result, arg[ part1.start ] )
                part1.start++
            }

            break
        }

        cmp, err := compare( arg[ part1.start ], arg[ part2.start ] )
        if err != nil {
            return fmt.Errorf( "comparison failed with error %v", err )
        }

        if cmp > 0 {
            result = append( result, arg[ part2.start ] )
            part2.start++
        } else {
            result = append( result, arg[ part1.start ] )
            part1.start++
        }
    }

    for _, elem := range result {
        arg[ i ] = elem
        i++
    }

    return nil
}

func MSort( arg interface{ }, compare Comparator )( result [ ]interface{ }, err error ) {
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

    li := len( result ) - 1

    for curSize := 1; curSize <= li; curSize *= 2 {
        for start := 0; start < li; start += 2 * curSize {
            mid := start + ( curSize - 1 )
            if mid > li {
                mid = li
            }

            end := start + ( ( 2 * curSize ) - 1 )
            if end > li {
                end = li
            }

            err := mergeParts( result, compare, part_t{ start : start, end : mid }, part_t{ start : mid + 1, end : end } )
            if nil != err {
                return nil, err
            }
        }
    }

    return result, nil
}
