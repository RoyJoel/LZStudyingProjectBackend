package utils

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

type IntMatrix3 [][][]int64

func (m *IntMatrix3) Scan(value interface{}) error {
	// 将读取到的字符串转换为 []byte 类型
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("IntMatrix.Scan: expected []byte got %T", value)
	}

	// 用 | 分割字符串，并将每个元素解析为 int 类型，存储到二维切片中
	str := string(bytes)
	rows := strings.Split(str, ";")
	matrix := make([][][]int64, len(rows))

	for i, row := range rows {
		xs := strings.Split(row, ",")
		matrix[i] = make([][]int64, len(xs))
		for j, x := range xs {
			cols := strings.Split(x, ".")
			matrix[i][j] = make([]int64, len(cols))
			for k, col := range cols {
				num, err := strconv.ParseInt(col, 10, 64)
				if err != nil {
					return fmt.Errorf("IntMatrix.Scan: failed to unmarshal data: %v", err)
				}
				matrix[i][j][k] = num
			}
		}

	}

	// 将解析得到的二维切片赋值给目标变量
	*m = matrix
	return nil
}

func (m IntMatrix3) Value() (driver.Value, error) {
	var buf bytes.Buffer

	for i, row := range m {
		for j, xs := range row {
			for k, x := range xs {
				if k != 0 {
					buf.WriteRune('.')
				}
				fmt.Fprintf(&buf, "%d", x)
			}
			if j != len(row)-1 {
				buf.WriteRune(',')
			}
		}
		if i != len(m)-1 {
			buf.WriteRune(';')
		}
	}

	return buf.Bytes(), nil
}

func (m IntMatrix3) Equals(other IntMatrix3) bool {
	if len(m) != len(other) {
		return false
	}
	for i, row := range m {
		if len(row) != len(other[i]) {
			return false
		}
		for j, col := range row {
			if len(col) != len(other[i][j]) {
				return false
			}
			for k, val := range col {
				if val != other[i][j][k] {
					return false
				}
			}
		}
	}
	return true
}

type IntMatrix2 [][]int64

func (m *IntMatrix2) Scan(value interface{}) error {
	// 将读取到的字符串转换为 []byte 类型
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("IntMatrix.Scan: expected []byte got %T", value)
	}

	// 用 | 分割字符串，并将每个元素解析为 int 类型，存储到二维切片中
	str := string(bytes)
	rows := strings.Split(str, ";")
	matrix := make([][]int64, len(rows))

	for i, row := range rows {
		xs := strings.Split(row, ".")
		matrix[i] = make([]int64, len(xs))
		for j, x := range xs {
			num, err := strconv.ParseInt(x, 10, 64)
			if err != nil {
				return fmt.Errorf("IntMatrix.Scan: failed to unmarshal data: %v", err)
			}
			matrix[i][j] = num

		}

	}

	// 将解析得到的二维切片赋值给目标变量
	*m = matrix
	return nil
}

func (m IntMatrix2) Value() (driver.Value, error) {
	var buf bytes.Buffer

	for i, row := range m {
		for k, x := range row {
			if k != 0 {
				buf.WriteRune('.')
			}
			fmt.Fprintf(&buf, "%d", x)
		}
		if i != len(m)-1 {
			buf.WriteRune(';')
		}
	}

	return buf.Bytes(), nil
}
func (m IntMatrix2) Equals(other IntMatrix2) bool {
	if len(m) != len(other) {
		return false
	}
	for i, row := range m {
		if len(row) != len(other[i]) {
			return false
		}
		for j, val := range row {
			if val != other[i][j] {
				return false
			}
		}
	}
	return true
}

type IntMatrix []int64

func (m *IntMatrix) Scan(value interface{}) error {
	// 将读取到的字符串转换为 []byte 类型
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("IntMatrix.Scan: expected []byte got %T", value)
	}

	// 用 | 分割字符串，并将每个元素解析为 int 类型，存储到二维切片中
	str := string(bytes)
	rows := strings.Split(str, ";")
	matrix := make([]int64, len(rows))

	for i, x := range rows {
		num, err := strconv.ParseInt(x, 10, 64)
		if err != nil {
			return fmt.Errorf("IntMatrix.Scan: failed to unmarshal data: %v", err)
		}
		matrix[i] = num

	}

	// 将解析得到的二维切片赋值给目标变量
	*m = matrix
	return nil
}

func (m IntMatrix) Value() (driver.Value, error) {
	var buf bytes.Buffer

	for i, row := range m {
		if i != 0 {
			buf.WriteRune(';')
		}
		fmt.Fprintf(&buf, "%d", row)
	}

	return buf.Bytes(), nil
}

func (m IntMatrix) Equals(other IntMatrix) bool {
	if len(m) != len(other) {
		return false
	}
	for i, val := range m {
		if val != other[i] {
			return false
		}
	}
	return true
}
