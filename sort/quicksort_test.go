package sort

import (
    "fmt"
    "testing"
    "time"
    "math/rand"
)

func comparator( a interface{ }, b interface{ } )( result int8, err error ) {
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

func testQsortOnce( t *testing.T ) {
    max_elems := rand.Intn( 256 )

    values := [ ]int{ }

    for i := 0; i < max_elems; i++ {
        values = append( values, rand.Intn( 256 ) )
    }

    result, err := Qsort( values, comparator )
    if nil != err {
        t.Errorf( "failed to sort" )
    }

    for i := 0; i < ( len( result ) - 1 ); i++ {
        if result[ i ].( int ) > result[ i + 1 ].( int ) {
            t.Logf( "%+v\n", values )
            t.Errorf( "incorrect sort" )
            break
        }
    }
}

func TestQsort( t *testing.T ) {
    rand.Seed( time.Now( ).UnixNano( ) )

    iters := rand.Intn( 16 )

    for i := 0; i < iters; i++ {
        testQsortOnce( t )
    }
}
