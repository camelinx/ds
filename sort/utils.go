package sort

import (
    "reflect"
)

func convertArgToSlice( arg interface{ } )( val reflect.Value, ok bool ) {
    val = reflect.ValueOf( arg )
    if reflect.Slice == val.Kind( ) {
        ok = true
    }

    return val, ok
}
