package chessboard

type File []bool
type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
	if file < "A" || file > "H" {
		return 0
	}

	c := 0
	for _, v := range cb[file] {
		if v {
			c++
		}
	}

	return c
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}

	c := 0
	for _, f := range cb {
		if f[rank-1] {
			c++
		}
	}

	return c
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	for _, f := range cb {
		return len(cb) * len(f)
	}

	return 0
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	c := 0
	for f := range cb {
		c += CountInFile(cb, f)
	}

	return c
}
