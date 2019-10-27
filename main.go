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

func citireListaArce() ([]int, []int) {
	file, err := ioutil.ReadFile("listaArce.txt")

	if err != nil {
		fmt.Println("Eroare citire matrice de incidenta.")
	}

	e1, e2 := []int{}, []int{}

	lines := strings.Split(string(file), "\r")
	var val int

	for i, line := range lines {
		vals := strings.Split(strings.TrimPrefix(line, "\n"), " ")

		for _, value := range vals {

			val, _ = strconv.Atoi(value)
			if i == 0 {
				e1 = append(e1, val)
			} else {
				e2 = append(e2, val)
			}
		}
	}

	return e1, e2
}

func citireListaSuc() ([]int, []int) {
	file, err := ioutil.ReadFile("listaSuccesori.txt")

	if err != nil {
		fmt.Println("Eroare citire matrice de incidenta.")
	}

	e1, e2 := []int{}, []int{}

	lines := strings.Split(string(file), "\r")
	var val int

	for i, line := range lines {
		vals := strings.Split(strings.TrimPrefix(line, "\n"), " ")

		for _, value := range vals {

			val, _ = strconv.Atoi(value)
			if i == 0 {
				e1 = append(e1, val)
			} else {
				e2 = append(e2, val)
			}
		}
	}

	return e1, e2
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

func incidentaToAdiacenta(matrix [][]int) [][]int {
	n := len(matrix)
	m := len(matrix[0])
	matrixA := make([][]int, n)
	for i := 0; i < n; i++ {
		matrixA[i] = make([]int, n)
	}

	var stIndex, enIndex int

	for i := 0; i < m; i++ {
		stIndex = -1
		enIndex = -1
		for j := 0; j < n; j++ {
			if matrix[j][i] == 1 {
				stIndex = j
			} else if matrix[j][i] == -1 {
				enIndex = j
			}

			if stIndex != -1 && enIndex != -1 {
				break
			}
		}

		matrixA[stIndex][enIndex] = 1
	}

	return matrixA
}

func incidentaToListaArce(matrix [][]int) ([]int, []int) {
	e1, e2 := []int{}, []int{}
	var stIndex, enIndex int

	n := len(matrix)
	m := len(matrix[0])

	for i := 0; i < m; i++ {
		stIndex = -1
		enIndex = -1
		for j := 0; j < n; j++ {
			if matrix[j][i] == 1 {
				stIndex = j
			} else if matrix[j][i] == -1 {
				enIndex = j
			}

			if stIndex != -1 && enIndex != -1 {
				e1 = append(e1, stIndex+1)
				e2 = append(e2, enIndex+1)
				break
			}
		}
	}

	return e1, e2
}

func incidentaToListaSuc(matrix [][]int) ([]int, []int) {
	e1, e2 := []int{}, []int{}
	var found bool

	n := len(matrix)
	m := len(matrix[0])
	poz := 1

	for i := 0; i < n; i++ {
		found = false
		e1 = append(e1, poz)

		for j := 0; j < m; j++ {
			if matrix[i][j] == 1 {
				for k := 0; k < n; k++ {
					if matrix[k][j] == -1 {
						poz++
						e2 = append(e2, k+1)
						found = true
						break
					}
				}
			}
		}

		if !found {
			poz++
		}

	}

	e1 = append(e1, poz-1)

	return e1, e2
}

func incidentaToListaPre(matrix [][]int) ([]int, []int) {
	e1, e2 := []int{}, []int{}
	var found bool

	n := len(matrix)
	m := len(matrix[0])
	poz := 1

	for i := 0; i < n; i++ {
		found = false
		e1 = append(e1, poz)

		for j := 0; j < m; j++ {
			if matrix[i][j] == -1 {
				for k := 0; k < n; k++ {
					if matrix[k][j] == 1 {
						poz++
						e2 = append(e2, k+1)
						found = true
						break
					}
				}
			}
		}

		if !found {
			poz++
		}

	}

	e1 = append(e1, poz-1)

	return e1, e2
}

func getNrNoduri(listaArce []int) int {
	if len(listaArce) == 0 {
		return 0
	}

	nr := 1
	lastNode := -1
	for _, val := range listaArce {
		if lastNode != val {
			lastNode = val
			nr++
		}
	}

	return nr
}

func listaArceToAdiacenta(v []int, w []int) [][]int {
	n := getNrNoduri(v)

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < len(v); i++ {
		matrix[v[i]-1][w[i]-1] = 1
	}

	return matrix
}

func listaArceToIncidenta(v []int, w []int) [][]int {
	n := getNrNoduri(v)
	m := len(v)

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
	}

	lastNode := v[0] - 1
	poz := 0
	for i := 0; i < len(v); i++ {
		if lastNode != v[i]-1 {
			lastNode = v[i] - 1
			poz++
		}

		matrix[poz][i] = 1
		matrix[w[i]-1][i] = -1
	}

	return matrix
}

func listaArceToListaSuc(v []int, w []int) ([]int, []int) {
	e1 := []int{}
	n := getNrNoduri(v)
	var found bool

	poz := 1
	j := 0

	for i := 0; i < n; i++ {
		e1 = append(e1, poz)
		found = false
		for j < len(w) && v[j]-1 == i {
			found = true
			j++
			poz++
		}

		if !found {
			poz++
		}
	}

	e1 = append(e1, poz-1)

	return e1, w
}

func listaArceToListaPre(v []int, w []int) ([]int, []int) {
	e1 := []int{}
	e2 := []int{}
	n := getNrNoduri(v)
	var found bool

	poz := 1

	for i := 0; i < n; i++ {
		e1 = append(e1, poz)
		found = false
		for j := 0; j < len(v); j++ {
			if w[j]-1 == i {
				found = true
				e2 = append(e2, v[j])
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

	/*	matrix := citireMatriceIncidenta()
		fmt.Println(matrix)

		matrixI := incidentaToAdiacenta(matrix)
		fmt.Println(matrixI)

		e1, e2 := incidentaToListaArce(matrix)
		fmt.Println(e1)
		fmt.Println(e2)

		s1, s2 := incidentaToListaSuc(matrix)
		fmt.Println(s1)
		fmt.Println(s2)

		p1, p2 := incidentaToListaPre(matrix)
		fmt.Println(p1)
		fmt.Println(p2)*/

	/*	e1, e2 := citireListaArce()
		fmt.Println(e1)
		fmt.Println(e2)

		matrix := listaArceToAdiacenta(e1, e2)
		fmt.Println(matrix)

		matrixI := listaArceToIncidenta(e1, e2)
		fmt.Println(matrixI)

		s1, s2 := listaArceToListaSuc(e1, e2)
		fmt.Println(s1)
		fmt.Println(s2)

		p1, p2 := listaArceToListaPre(e1, e2)
		fmt.Println(p1)
		fmt.Println(p2)*/

	e1, e2 := citireListaSuc()
	fmt.Println(e1)
	fmt.Println(e2)
}
