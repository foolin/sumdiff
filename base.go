package sumdiff

import (
	"github.com/foolin/sumdiff/internal/statusbar"
	"github.com/foolin/sumdiff/internal/util"
)

func listPathWithStatusbar(path string) (map[string]util.PathInfo, error) {
	return util.ListPath(path, func(info util.PathInfo) bool {
		statusbar.Display("Listing %v", info.Path)
		return true
	})
}
