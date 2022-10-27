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
