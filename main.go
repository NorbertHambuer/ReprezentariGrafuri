package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
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

func citireListaArce() (int, []int, []int) {
	file, err := ioutil.ReadFile("listaArce.txt")

	if err != nil {
		fmt.Println("Eroare citire matrice de incidenta.")
	}

	e1, e2 := []int{}, []int{}

	lines := strings.Split(string(file), "\r")
	var val, n int

	for i, line := range lines {
		vals := strings.Split(strings.TrimPrefix(line, "\n"), " ")

		if len(vals) == 1 {
			val, _ = strconv.Atoi(vals[0])
			n = val
			continue
		}

		for _, value := range vals {

			val, _ = strconv.Atoi(value)
			if i == 1 {
				e1 = append(e1, val)
			} else {
				e2 = append(e2, val)
			}
		}
	}

	return n, e1, e2
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

func citireListaPre() ([]int, []int) {
	file, err := ioutil.ReadFile("listaPredecesori.txt")

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

func listaArceToAdiacenta(n int, v []int, w []int) [][]int {
	//n := getNrNoduri(v)

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < len(v); i++ {
		matrix[v[i]-1][w[i]-1] = 1
	}

	return matrix
}

func listaArceToIncidenta(n int, v []int, w []int) [][]int {
	//n := getNrNoduri(v)
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

func listaArceToListaSuc(n int, v []int, w []int) ([]int, []int) {
	e1 := []int{}
	//n := getNrNoduri(v)
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

func listaArceToListaPre(n int, v []int, w []int) ([]int, []int) {
	e1 := []int{}
	e2 := []int{}
	//n := getNrNoduri(v)
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

func listaSucToAdiacenta(v []int, w []int) [][]int {
	n := len(v) - 1
	m := len(w) - 1

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := v[i]; j < v[i+1]; j++ {
			matrix[i][w[j-1]-1] = 1
		}
	}

	if v[n] != v[n-1] {
		matrix[n-1][w[m]-1] = 1
	}

	return matrix
}

func listaSucToIncidenta(v []int, w []int) [][]int {
	n := len(v) - 1
	m := len(w)

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
	}

	poz := 0
	for i := 0; i < n; i++ {
		for j := v[i]; j < v[i+1]; j++ {
			matrix[i][poz] = 1
			matrix[w[j-1]-1][poz] = -1
			poz++
		}
	}

	if v[n] != v[n-1] {
		matrix[w[m-1]-1][poz] = -1
		matrix[n-1][poz] = 1
	}

	return matrix
}

func listaSucToListaArce(v []int, w []int) ([]int, []int) {
	e1 := []int{}
	e2 := []int{}

	n := len(v) - 1

	for i := 0; i < n; i++ {
		for j := v[i]; j < v[i+1]; j++ {
			e1 = append(e1, i+1)
			e2 = append(e2, w[j-1])
		}
	}

	if v[n] != v[n-1] {
		m := len(w)

		e1 = append(e1, n)
		e2 = append(e2, w[m-1])
	}

	return e1, e2
}

func listaSucToListaPre(v []int, w []int) ([]int, []int) {
	e1 := []int{}
	e2 := []int{}
	n := len(v) - 1
	var found bool
	var node, nodeIndex int

	poz := 1

	for i := 0; i < n; i++ {
		e1 = append(e1, poz)
		found = false

		for j := 0; j < len(w); j++ {
			if w[j]-1 == i {
				found = true

				node = 0
				nodeIndex = v[0]

				for nodeIndex-1 <= j && node < n {
					node++
					nodeIndex = v[node]
				}

				e2 = append(e2, node)
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

func listaPreToAdiacenta(v []int, w []int) [][]int {
	n := len(v) - 1
	m := len(w)

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {

		for j := v[i]; j < v[i+1]; j++ {
			matrix[w[j-1]-1][i] = 1
		}
	}

	if v[n] != v[n-1] {
		matrix[w[m-1]-1][n-1] = 1
	}

	return matrix
}

func listaPreToIncidenta(v []int, w []int) [][]int {
	n := len(v) - 1
	m := len(w)

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
	}

	var node, nodeIndex int
	poz := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if w[j]-1 == i {
				matrix[i][poz] = 1

				node = 0
				nodeIndex = v[0]

				for nodeIndex-1 <= j && node < n {
					node++
					nodeIndex = v[node]
				}

				matrix[node-1][poz] = -1

				poz++
			}
		}
	}

	return matrix
}

func listaPreToListaArce(v []int, w []int) ([]int, []int) {
	e1 := []int{}
	e2 := []int{}

	n := len(v) - 1
	m := len(w)

	var node, nodeIndex int
	poz := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if w[j]-1 == i {

				e1 = append(e1, i+1)

				node = 0
				nodeIndex = v[0]

				for nodeIndex-1 <= j && node < n {
					node++
					nodeIndex = v[node]
				}

				e2 = append(e2, node)

				poz++
			}
		}
	}

	return e1, e2
}

func listaPreToListaSuc(v []int, w []int) ([]int, []int) {
	e1 := []int{}
	e2 := []int{}

	n := len(v) - 1
	m := len(w)

	var node, nodeIndex int
	var found bool
	poz := 1

	for i := 0; i < n; i++ {

		e1 = append(e1, poz)
		found = false
		for j := 0; j < m; j++ {
			if w[j]-1 == i {
				found = true

				node = 0
				nodeIndex = v[0]

				for nodeIndex-1 <= j && node < n {
					node++
					nodeIndex = v[node]
				}

				e2 = append(e2, node)

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

func printMeniu() {
	fmt.Println("1 - Transformare matrice adiacenta")
	fmt.Println("2 - Transformare matrice incidenta")
	fmt.Println("3 - Transformare lista arce")
	fmt.Println("4 - Transformare lista succesori")
	fmt.Println("5 - Transformare lista predecesor")
	fmt.Println("6 - Inchidere")
}

func main() {
	end := false
	reader := bufio.NewReader(os.Stdin)
	for !end {
		printMeniu()
		text, _ := reader.ReadString('\n')

		cheie, err := strconv.Atoi(strings.TrimSpace(text))

		if err != nil || cheie < 0 || cheie > 6 {
			fmt.Println("Comanda invalida")
		}

		switch cheie {
		case 1:
			matrix := citireMatriceAdiacenta()
			fmt.Println(matrix)

			matrixI := adiacentaToIncidenta(matrix)
			fmt.Println(matrixI)

			e1, e2 := adiacentaToListaArce(matrix)
			fmt.Println(e1)
			fmt.Println(e2)

			s1, s2 := adiacentaToListaSuc(matrix)
			fmt.Println(s1)
			fmt.Println(s2)

			p1, p2 := adiacentaToListaPre(matrix)
			fmt.Println(p1)
			fmt.Println(p2)
		case 2:
			matrix := citireMatriceIncidenta()
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
			fmt.Println(p2)
		case 3:
			n, e1, e2 := citireListaArce()
			fmt.Println(n)
			fmt.Println(e1)
			fmt.Println(e2)

			matrix := listaArceToAdiacenta(n, e1, e2)
			fmt.Println(matrix)

			matrixI := listaArceToIncidenta(n, e1, e2)
			fmt.Println(matrixI)

			s1, s2 := listaArceToListaSuc(n, e1, e2)
			fmt.Println(s1)
			fmt.Println(s2)

			p1, p2 := listaArceToListaPre(n, e1, e2)
			fmt.Println(p1)
			fmt.Println(p2)
		case 4:
			e1, e2 := citireListaSuc()
			fmt.Println(e1)
			fmt.Println(e2)

			matrix := listaSucToAdiacenta(e1, e2)
			fmt.Println(matrix)

			matrixI := listaSucToIncidenta(e1, e2)
			fmt.Println(matrixI)

			s1, s2 := listaSucToListaArce(e1, e2)
			fmt.Println(s1)
			fmt.Println(s2)

			p1, p2 := listaSucToListaPre(e1, e2)
			fmt.Println(p1)
			fmt.Println(p2)
		case 5:
			e1, e2 := citireListaPre()
			fmt.Println(e1)
			fmt.Println(e2)

			matrix := listaPreToAdiacenta(e1, e2)
			fmt.Println(matrix)

			matrixI := listaPreToIncidenta(e1, e2)
			fmt.Println(matrixI)

			s1, s2 := listaPreToListaArce(e1, e2)
			fmt.Println(s1)
			fmt.Println(s2)

			p1, p2 := listaPreToListaSuc(e1, e2)
			fmt.Println(p1)
			fmt.Println(p2)
		case 6:
			end = true
		}
		fmt.Println()
	}
}
