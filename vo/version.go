package vo

import (
	"fmt"
	"runtime"
	"strings"
	"text/tabwriter"
)

type AppInfo struct {
	Description string `json:"description" yaml:"description"`
	Version     string `json:"version" yaml:"version"`
	Commit      string `json:"commit" yaml:"commit"`
	Date        string `json:"date" yaml:"date"`
	Compiler    string `json:"compiler" yaml:"compiler"`
	Platform    string `json:"platform" yaml:"platform"`
}

func NewAppInfo() AppInfo {
	return AppInfo{
		Description: "",
		Version:     "devel",
		Commit:      "none",
		Date:        "none",
		Compiler:    runtime.Version(),
		Platform:    fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

func (r AppInfo) String() string {
	b := strings.Builder{}
	w := tabwriter.NewWriter(&b, 0, 0, 2, ' ', 0)
	if r.Description != "" {
		_, _ = fmt.Fprint(w, r.Description)
		_, _ = fmt.Fprint(w, "\n\n")
	}
	_, _ = fmt.Fprintf(w, "Version:\t%s\n", r.Version)
	_, _ = fmt.Fprintf(w, "Commit:\t%s\n", r.Commit)
	_, _ = fmt.Fprintf(w, "Date:\t%s\n", r.Date)
	_, _ = fmt.Fprintf(w, "Compiler:\t%s\n", r.Compiler)
	_, _ = fmt.Fprintf(w, "Platform:\t%s\n", r.Platform)

	_ = w.Flush()
	return b.String()
}