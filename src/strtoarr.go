package main 

func count_rows(str []byte) int {
	rows := 0

	for i := 0; i < len(str); i++ {
		if str[i] == '\n' {
			rows++
		}
	}
	return rows
}

func count_cols(str []byte, rows int) []int {
    cols := make([]int, rows)
    x := 0

    for y := 0; y < rows; y++ {
        for ; str[x] != '\n'; x++ {
            cols[y]++;
        }
        x++
    }
    return cols
}

func str_to_arr(str []byte) [][]byte {
    rows := count_rows(str)
    cols := count_cols(str, rows)
    arr := make([][]byte, rows)
    x := 0

    for y := 0; y < rows; y++ {
        arr[y] = make([]byte, cols[y])
        for w := 0; str[x] != '\n'; x++ {
            arr[y][w] = str[x]
            w++
        }
        x++
    }
    return arr
}
