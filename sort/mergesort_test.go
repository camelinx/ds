package sort

import (
    "fmt"
    "testing"
    "time"
    "math/rand"
)

func msort_comparator( a interface{ }, b interface{ } )( result int8, err error ) {
    aval, ok := a.( int )
    if !ok {
        return 0, fmt.Errorf( "unsupported value type" )
    }

    bval, ok := b.( int )
    if !ok {
        return 0, fmt.Errorf( "unsupported value type" )
    }

    if aval == bval {
        return 0, nil
    } else if aval > bval {
        return 1, nil
    }

    return -1, nil
}

func testMSortOnce( t *testing.T ) {
    max_elems := rand.Intn( 256 )

    values := [ ]int{ }
    lookup := make( map[ int ]bool )

    for i := 0; i < max_elems; i++ {
        value  := rand.Intn( 256 )
        values  = append( values, value )

        lookup[ value ] = true
    }

    result, err := MSort( values, msort_comparator )
    if nil != err {
        t.Errorf( "failed to sort" )
        return
    }

    if len( result ) != len( values ) {
        t.Errorf( "incorrect sort - length mismatch" )
        return
    }

    for i := 0; i < ( len( result ) - 1 ); i++ {
        if _, exists := lookup[ result[ i ].( int ) ]; !exists {
            t.Errorf( "incorrect sort - unexpected value in returned list" )
            break
        }

        if result[ i ].( int ) > result[ i + 1 ].( int ) {
            t.Errorf( "incorrect sort" )
            break
        }
    }
}

func TestMSort( t *testing.T ) {
    rand.Seed( time.Now( ).UnixNano( ) )

    iters := rand.Intn( 16 )

    for i := 0; i < iters; i++ {
        testMSortOnce( t )
    }
}
