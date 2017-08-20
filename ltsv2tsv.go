package ltsv2tsv

import (
	"bufio"
	"io"
	"strings"
)

type LTSVConverter struct {
	reader *bufio.Reader
}

func NewConverter(r io.Reader) *LTSVConverter {
	return &LTSVConverter{bufio.NewReader(r)}
}

func (r *LTSVConverter) Converter() (records [][]string, err error) {
	header := map[string]int{}
	datas := map[int][][]string{}
	data_index := 0

	// create header record and parse data records.
	for {
		line, _, err := r.reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		str_line := strings.TrimSpace(string(line))
		if str_line == "" {
			continue
		}

		line_arr := strings.Split(str_line, "\t")
		fields := make([][]string, len(line_arr))

		for i, field := range line_arr {
			field_arr := strings.Split(field, ":")

			_, is_exists := header[field_arr[0]]

			if !is_exists {
				header[field_arr[0]] = len(header)
			}
			fields[i] = field_arr
		}
		datas[data_index] = fields
		data_index += 1
	}

	head_record := make([]string, len(header))
	// from: {"header3": 2, "header1": 0, "header2": 1}
	// to  : ["header1", "header2", "header3"]
	for k, v := range header {
		head_record[v] = k
	}

	records = make([][]string, data_index+1)

	for i := 0; i < len(datas); i++ {
		if i == 0 {
			records[i] = head_record
		}
		record := make([]string, len(header))
		for _, fields := range datas[i] {
			pos, _ := header[fields[0]]
			record[pos] = fields[1]
		}
		records[i+1] = record
	}

	return
}
