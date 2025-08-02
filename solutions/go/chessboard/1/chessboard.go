package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) (count int) {
	for fileKey, fileValue := range cb {
        for _, occupied := range fileValue {
            if occupied && fileKey == file {
                count++
            }
        }
    }

    return
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) (count int) {
	for _, file := range cb {
        for currentRank, occupied := range file {
            if occupied && rank == currentRank + 1 {
                count++
            }
        }
    }

    return
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) (count int) {
		for _, file := range cb {
	        for range file {
    	        count++
        	}
    	}

    return
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) (count int) {
	for _, file := range cb {
        for _, occupied := range file {
            if occupied {
                count++
            }
        }
    }

    return
}
