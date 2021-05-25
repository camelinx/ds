package bst

import (
    "fmt"
    "testing"
    "time"
    "math/rand"
)

func TestInit( t *testing.T ) {
    tree := Init( )
    if nil != tree.Root || 0 != tree.Count {
        t.Errorf( "tree returned by init %+v not initialized", tree )
    }
}

func comparator( a interface{ }, b interface{ } )( result int8, err error ) {
    aval, ok := a.( uint )
    if !ok {
        return 0, fmt.Errorf( "unsupported value type" )
    }

    bval, ok := b.( uint )
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

const (
    max_tree_val   = 256
    max_tree_nodes = 31
)

func ( tree *Bst_t )insertToTree( t *testing.T, withDup bool )( values map[ uint ]bool ) {
    rand.Seed( time.Now( ).UnixNano( ) )

    values = make( map[ uint ]bool )

    for i := 0; i < max_tree_nodes; i++ {
        value := uint( rand.Intn( max_tree_val ) )
        if _, exists := values[ value ]; exists {
            new_value := value + 1
            for new_value != value {
                if _, exists = values[ new_value ]; !exists {
                    break
                }

                if max_tree_val == new_value {
                    new_value = 0
                } else {
                    new_value++
                }
            }

            if new_value == value {
                t.Errorf( "failed to find unique value to insert into tree %+v", tree )
            }

            value = new_value
        }

        values[ value ] = true

        count, err := tree.Insert( value, comparator )
        if nil != err || uint( i + 1 ) != count || count != tree.Count {
            t.Errorf( "failed to insert into tree %+v", tree )
        }

        if withDup {
            count, err = tree.Insert( value, comparator )
            if nil != err || uint( i + 1 ) != count || count != tree.Count {
                t.Errorf( "failed to insert (duplicate) into tree %+v", tree )
            }
        }
    }

    return values
}

func TestInsert( t *testing.T ) {
    tree := Init( )
    if nil == tree || nil != tree.Root || 0 != tree.Count {
        t.Errorf( "tree returned by init %+v not initialized", tree )
    }

    tree.insertToTree( t, true )

    _, err := tree.Insert( "hello", comparator )
    if nil == err {
        t.Errorf( "successfully inserted a string into a binary search tree %+v of unsigned integers", tree )
    }

    _, err = tree.Insert( 10, comparator )
    if nil == err {
        t.Errorf( "successfully inserted an integer into a binary search tree %+v of unsigned integers", tree )
    }

    _, err = tree.Insert( 1.0, comparator )
    if nil == err {
        t.Errorf( "successfully inserted a float into a binary search tree %+v of unsigned integers", tree )
    }

    tree = nil
    _, err = tree.Insert( uint( 10 ), comparator )
    if nil == err {
        t.Errorf( "should panic before we get here: attempt to insert into nil tree" )
    }
}

func( tree *Bst_t )searchInTree( t *testing.T, values map[ uint ]bool ) {
    for i := uint( 0 ); i < max_tree_val; i++ {
        found, err := tree.Search( i, comparator )
        if err != nil {
            t.Errorf( "error %v while searching %v in tree %+v", err, i, tree )
        }

        if _, exists := values[ i ]; exists {
            if !found {
                t.Errorf( "failed to find value %v in tree %+v", i, tree )
            }
        } else {
            if found {
                t.Errorf( "found non existent value %v in tree %+v", i, tree )
            }
        }
    }
}

func TestSearch( t *testing.T ) {
    tree := Init( )
    if nil == tree || nil != tree.Root || 0 != tree.Count {
        t.Errorf( "tree returned by init %+v not initialized", tree )
    }

    values := tree.insertToTree( t, false )

    tree.searchInTree( t, values )

    _, err := tree.Search( "hello", comparator )
    if err == nil {
        t.Errorf( "successfully searched a string in a tree %+v of unsigned integers", tree )
    }

    _, err = tree.Search( 1.0, comparator )
    if err == nil {
        t.Errorf( "successfully searched a float in a tree %+v of unsigned integers", tree )
    }

    tree = nil
    _, err = tree.Search( uint( 10 ), comparator )
    if nil == err {
        t.Errorf( "should panic before we get here: attempt to insert into nil tree" )
    }
}

func ( tree *Bst_t )deleteFromTree( t *testing.T, values map[ uint ]bool ) {
    fCount := tree.Count

    for i := uint( 0 ); i < max_tree_val; i++ {
        count, err := tree.Delete( i, comparator )
        if err != nil && tree.Count > 0 {
            t.Errorf( "error %v while deletng %v from tree %+v", err, i, tree )
        }

        if _, exists := values[ i ]; exists {
            if ( fCount - count ) != 1 || count != tree.Count {
                t.Errorf( "failed to delete value %v from tree %+v", i, tree )
            }

            fCount--
        } else {
            if fCount != count || count != tree.Count {
                t.Errorf( "deleted non existent value %v from tree %+v", i, tree )
            }
        }
    }
}

func TestDelete( t *testing.T ) {
    tree := Init( )
    if nil == tree || nil != tree.Root || 0 != tree.Count {
        t.Errorf( "tree returned by init %+v not initialized", tree )
    }

    values := tree.insertToTree( t, false )

    tree.deleteFromTree( t, values )

    _, err := tree.Delete( "hello", comparator )
    if err == nil {
        t.Errorf( "successfully deleted a string from tree %+v of unsigned integers", tree )
    }

    _, err = tree.Delete( 1.0, comparator )
    if err == nil {
        t.Errorf( "successfully deleted a float from tree %+v of unsigned integers", tree )
    }

    tree = nil
    _, err = tree.Delete( uint( 10 ), comparator )
    if nil == err {
        t.Errorf( "should panic before we get here: attempt to delete from a nil tree" )
    }
}
