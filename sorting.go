
package algorithms

func Bubble_sort_modified(input []int) {
	n := len(input) //n er antallet i lista som skal bli sortert
	swapped := true
		for swapped{ //loop
		swapped = false
			for i := 1; i < n; i++ { //går gjennom hvert tall i lista
				if input[i-1] > input[i] { // hvis nårværende element > neste = swap
					input[i], input[i-1] = input[i-1], input[i]
					swapped = true



				}


			}


		}

}




func Bubble_sort(list []int) []int {
	// find the length of list n
	n := len(list)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j+1]
				list[j+1] = list[j]
				list[j] = temp
			}
		}
	}

	return list
}

// Implementering av Quicksort algoritmen
func QSort(values []int) {
	qsort(values, 0, len(values)-1)
}

func qsort(values []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := values[l]
	i := l + 1

	for j := l; j <= r; j++ {
		if pivot > values[j] {
			values[i], values[j] = values[j], values[i]
			i++
		}
	}

	values[l], values[i-1] = values[i-1], pivot

	qsort(values, l, i-2)
	qsort(values, i, r)
}
