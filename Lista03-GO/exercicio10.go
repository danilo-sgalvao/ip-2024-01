package main

import "fmt"

func main() {
	var N, maxNote, maxNoteIndex int
	fmt.Scan(&N)

	notes := make([]int, N)
	frequency := make(map[int]int)

	for i := 0; i < N; i++ {
		fmt.Scan(&notes[i])
		frequency[notes[i]] = i
		if notes[i] > maxNote {
			maxNote = notes[i]
			maxNoteIndex = i
		}
	}

	lastNote := notes[N-1]

	count := 0
	for _, v := range notes {
		if v == lastNote {
			count++
		}
	}

	fmt.Printf("Nota %d, %d vezes\n", lastNote, count)
	fmt.Printf("Nota %d, indice %d\n", maxNote, maxNoteIndex)
}
