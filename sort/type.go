package sort

type Comparator func( interface{ }, interface{ } )( int8, error )

type SortOrder int

const (
    Ascending   SortOrder   = 0

    Descending  SortOrder   = 1
)

type part_t struct {
    start   int
    end     int
}
