package write

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/liushuochen/gotable"
	"gopkg.in/yaml.v3"
	"io"
)

type Writer struct {
	w io.Writer
}

func (r *Writer) Printf(s string, a ...any) (n int, err error) {
	return r.w.Write([]byte(fmt.Sprintf(s, a...)))
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
	_, err = r.Printf("%v", table)
	if err != nil {
		return err
	}
	return nil
}

func (r *Writer) Csv(records [][]string) error {
	writer := csv.NewWriter(r.w)
	return writer.WriteAll(records)
}

func (r *Writer) Json(v any) error {
	data, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return err
	}
	_, err = r.Printf("%s", data)
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
	_, err = r.Printf("%s", data)
	if err != nil {
		return err
	}
	return nil
}
