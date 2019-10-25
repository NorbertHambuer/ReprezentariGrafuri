package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func citireMatriceAdiacenta() [][]int {
	file, err := ioutil.ReadFile("matriceAdiacenta.txt")

	if err != nil {
		fmt.Println("Eroare citire matrice de adiacenta.")
	}

	lines := strings.Split(string(file), "\r")
	n := len(lines)

	matrix := make([][]int, n)

	for i, line := range lines {
		matrix[i] = make([]int, n)

		vals := strings.Split(strings.TrimPrefix(line, "\n"), " ")

		for j, value := range vals {
			matrix[i][j], _ = strconv.Atoi(value)
		}
	}

	return matrix
}

func citireMatriceIncidenta() [][]int {
	file, err := ioutil.ReadFile("matriceIncidenta.txt")

	if err != nil {
		fmt.Println("Eroare citire matrice de incidenta.")
	}

	lines := strings.Split(string(file), "\r")
	n := len(lines)

	matrix := make([][]int, n)

	for i, line := range lines {
		vals := strings.Split(strings.TrimPrefix(line, "\n"), " ")

		matrix[i] = make([]int, len(vals))

		for j, value := range vals {
			matrix[i][j], _ = strconv.Atoi(value)
		}
	}

	return matrix
}

func adiacentaToIncidenta(matrix [][]int) [][]int {
	nodes := 0
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				nodes++
			}
		}
	}

	matrixI := make([][]int, n)
	nodeIndex := 0

	for i := 0; i < n; i++ {
		matrixI[i] = make([]int, nodes)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				matrixI[i][nodeIndex] = 1
				matrixI[j][nodeIndex] = -1
				nodeIndex++
			}
		}
	}

	return matrixI
}

func adiacentaToListaArce(matrix [][]int) ([]int, []int) {
	e1, e2 := []int{}, []int{}
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				e1 = append(e1, i+1)
				e2 = append(e2, j+1)
			}
		}
	}

	return e1, e2
}

func adiacentaToListaSuc(matrix [][]int) ([]int, []int) {
	e1, e2 := []int{}, []int{}
	poz := 1
	n := len(matrix)
	var found bool

	for i := 0; i < n; i++ {
		e1 = append(e1, poz)
		found = false

		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				e2 = append(e2, j+1)
				found = true
				poz++
			}
		}

		if !found {
			poz++
		}
	}

	e1 = append(e1, poz-1)

	return e1, e2
}

func adiacentaToListaPre(matrix [][]int) ([]int, []int) {
	e1, e2 := []int{}, []int{}
	poz := 1
	n := len(matrix)
	var found bool

	for i := 0; i < n; i++ {
		e1 = append(e1, poz)
		found = false

		for j := 0; j < n; j++ {
			if matrix[j][i] == 1 {
				e2 = append(e2, j+1)
				found = true
				poz++
			}
		}

		if !found {
			poz++
		}
	}

	e1 = append(e1, poz-1)

	return e1, e2
}

func IncidentaToAdiacenta(matrix [][]int) [][]int {

}

func main() {
	/*	matrix := citireMatriceAdiacenta()
		fmt.Println(matrix)

		matrixI := adiacentaToIncidenta(matrix)
		fmt.Println(matrixI)

		e1,e2 := adiacentaToListaArce(matrix)
		fmt.Println(e1)
		fmt.Println(e2)

		s1, s2 := adiacentaToListaSuc(matrix)
		fmt.Println(s1)
		fmt.Println(s2)

		p1, p2 := adiacentaToListaPre(matrix)
		fmt.Println(p1)
		fmt.Println(p2)*/

	matrix := citireMatriceIncidenta()
	fmt.Println(matrix)
}
