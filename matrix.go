package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix struct {
	values [][]int
}

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	matrix := Matrix{values: make([][]int, len(rows))}

	if len(rows) == 0 {
		return nil, nil
	}

	for i, r := range rows {
		values := strings.Split(r, " ")
		rowOk := false
		for _, c := range values {
			if c != "" && c != " " {
				v, err := strconv.Atoi(c)
				if err != nil {
					return nil, errors.New("Invalid Matrix!")
				} else {
					matrix.values[i] = append(matrix.values[i], v)
					rowOk = true
				}
			}
		}
		if !rowOk {
			return nil, errors.New("Invalid Matrix!")
		}
		if i > 0 && len(matrix.values[i]) != len(matrix.values[i-1]) {
			return nil, errors.New("Matrix can not have different columns size!")
		}
	}
	return &matrix, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	cols := make([][]int, len(m.values[0]))
	for i := 0; i < len(cols); i++ {
		cols[i] = make([]int, len(m.values))
		for j := 0; j < len(cols[i]); j++ {
			cols[i][j] = m.values[j][i]
		}
	}
	return cols
}

func (m *Matrix) Rows() [][]int {
	rows := make([][]int, len(m.values))
	for i := 0; i < len(m.values); i++ {
		rows[i] = make([]int, len(m.values[0]))
		for j := 0; j < len(m.values[0]); j++ {
			rows[i][j] = m.values[i][j]
		}
	}
	return rows
}

func (m *Matrix) Set(row, col, val int) bool {
	if row >= 0 && row < len(m.values) && col >= 0 && col < len(m.values[0]) {
		m.values[row][col] = val
		return true
	} else {
		return false
	}
}
