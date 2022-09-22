package utils

// GetTransformationDistance uses levenshtein distance algorithm
// to get the minimun distance to convert text1 into text2
func GetTransformationDistance(text1, text2 string) int {
	// 1. Get unicode characters for text1 and text2 strings using rune data type
	t1 := []rune(text1)
	t2 := []rune(text2)

	t1Len := len(t1)
	t2Len := len(t2)

	// 2. Create slice with length of text1 + 1
	column := make([]int, t1Len+1)

	// 3. Assign incremental values to column slice
	for y := 1; y <= t1Len; y++ {
		column[y] = y
	}

	// 3. Iterate the items of text 1 with the items of text 2
	for x := 1; x <= t2Len; x++ {
		// assign the last index position of text 2 in column[0] slice
		column[0] = x
		lastkey := x - 1 // last index position before the new assigment

		for y := 1; y <= t1Len; y++ {
			oldkey := column[y]
			var incr int

			// if there's a difference between a character of text 1 and a character of text 2

			if t1[y-1] != t2[x-1] {
				incr = 1
			}

			// gets the minimun value of insert, delete and remove operation
			// it is assumed that there will be at least one update operation after an insert, delete or replace
			column[y] = minimum(column[y]+1, column[y-1]+1, lastkey+incr)
			lastkey = oldkey
		}

	}
	return column[t1Len]
}

func minimum(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}
