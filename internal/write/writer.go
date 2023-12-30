package write

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/liushuochen/gotable"
	"gopkg.in/yaml.v3"
)

type Writer struct {
	writer io.Writer
	format Format
}

type TableData interface {
	Array() [][]string
}

func New(w io.Writer, f Format) *Writer {
	return &Writer{
		writer: w,
		format: f,
	}
}

func NewStd() *Writer {
	return New(os.Stdout, None)
}

func (r *Writer) Write(data TableData) error {
	switch r.format {
	case Table:
		return r.Table(data.Array())
	case Csv:
		return r.Csv(data.Array())
	case Json:
		return r.Json(data)
	case Yaml:
		return r.Yaml(data)
	case None:
		fallthrough
	default:
		return r.Table(data.Array())
	}
}

func (r *Writer) MustWrite(data TableData) {
	err := r.Write(data)
	if err != nil {
		fmt.Printf("Write error: %v", err)
		os.Exit(1)
	}
}

func (r *Writer) Printf(s string, a ...any) (n int, err error) {
	return r.writer.Write([]byte(fmt.Sprintf(s, a...)))
}

func (r *Writer) Table(records [][]string) error {
	if len(records) == 0 {
		return fmt.Errorf("table is empty")
	}
	table, err := gotable.Create(records[0]...)
	if err != nil {
		return err
	}
	for i, record := range records {
		if i == 0 {
			continue //Header
		}
		err = table.AddRow(record)
		if err != nil {
			return err
		}
	}
	_, err = r.Printf("%v\n", table)
	if err != nil {
		return err
	}
	return nil
}

func (r *Writer) Csv(records [][]string) error {
	writer := csv.NewWriter(r.writer)
	return writer.WriteAll(records)
}

func (r *Writer) Json(v any) error {
	data, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return err
	}
	_, err = r.Printf("%s\n", data)
	if err != nil {
		return err
	}
	return nil
}

func (r *Writer) Yaml(v any) error {
	data, err := yaml.Marshal(v)
	if err != nil {
		return err
	}
	_, err = r.Printf("%s\n", data)
	if err != nil {
		return err
	}
	return nil
}
