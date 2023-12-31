package sumdiff

import (
	"fmt"
	"os"

	"github.com/foolin/sumdiff/vo"

	"github.com/foolin/sumdiff/internal/statusbar"
	"github.com/foolin/sumdiff/internal/util"
	"github.com/hashicorp/go-multierror"
)

func Cmp(path1, path2 string) (ok bool, res vo.CmpList, err error) {
	statusbar.Display("Comparing %v <-> %v", path1, path2)

	path1 = util.FormatPath(path1)
	path2 = util.FormatPath(path2)

	result := vo.NewCmpInfo(path1, path2)
	result.Equal = false
	var outError error

	s1, err := os.Stat(path1)
	if err != nil {
		result.Hash1.Error = err
		outError = multierror.Append(outError, err)
	}

	s2, err := os.Stat(path2)
	if err != nil {
		result.Hash2.Error = err
		outError = multierror.Append(outError, err)
	}

	if outError != nil {
		result.Msg = outError.Error()
		return false, []vo.CmpInfo{result}, outError
	}

	if s1.IsDir() != s2.IsDir() {
		//result.Msg = fmt.Sprintf("file type not equal [ %v != %v ]", util.FileType(s1.IsDir()), util.FileType(s2.IsDir()))
		result.Msg = "not equal file type"
		return false, []vo.CmpInfo{result}, outError
	}

	if s1.IsDir() {
		return CmpDir(path1, path2)
	} else {
		ok, ret, err := CmpFile(path1, path2)
		return ok, []vo.CmpInfo{ret}, err
	}
}

func CmpDir(path1, path2 string) (bool, vo.CmpList, error) {
	var outError error
	outResult := vo.NewCmpInfo(path1, path2)

	data1, err := listPathWithStatusbar(path1)
	if err != nil {
		outResult.Hash1.Error = err
		outError = multierror.Append(outError, err)
	}

	data2, err := listPathWithStatusbar(path2)
	if err != nil {
		outResult.Hash2.Error = err
		outError = multierror.Append(outError, err)
	}

	if outError != nil {
		outResult.Equal = false
		outResult.Msg = outError.Error()
		return false, []vo.CmpInfo{outResult}, outError
	}

	notEqualCount := 0
	retList := make(vo.CmpList, 0)

	for k, v1 := range data1 {
		statusbar.Display("Comparing " + k)
		itemResult := vo.NewCmpInfo(k, k)
		itemResult.Hash1.Size = v1.Info.Size()
		itemResult.Equal = false
		v2, ok := data2[k]

		if !ok {
			itemResult.Hash2.Error = fmt.Errorf("not exist [%v]", k)
			//itemResult.Msg = fmt.Sprintf("path2 not exist [%v]", k)
			itemResult.Msg = "path2 not exist"
			retList = append(retList, itemResult)
			notEqualCount++
			continue
		} else {
			itemResult.Hash2.Size = v2.Info.Size()
		}

		if v1.Info.Size() != v2.Info.Size() {
			//itemResult.Msg = fmt.Sprintf("size not equal [ %v != %v ]", v1.Info.Size(), v2.Info.Size())
			itemResult.Msg = "not equal size"
			retList = append(retList, itemResult)
			notEqualCount++
			continue
		}

		if v1.Info.IsDir() != v2.Info.IsDir() {
			//itemResult.Msg = fmt.Sprintf("file type not equal [ %v != %v ]", util.FileType(v1.Info.IsDir()), util.FileType(v2.Info.IsDir()))
			itemResult.Msg = fmt.Sprintf("not equal file type")
			retList = append(retList, itemResult)
			notEqualCount++
			continue
		}

		//File check file md5
		if !v1.Info.IsDir() {
			var itemError error

			h1, err := util.Sha256(v1.Path)
			if err != nil {
				itemResult.Hash1.Error = err
				itemError = multierror.Append(outError, err)
				outError = multierror.Append(outError, err)
			} else {
				itemResult.Hash1.Hash = h1
			}

			h2, err := util.Sha256(v2.Path)
			if err != nil {
				itemResult.Hash2.Error = err
				itemError = multierror.Append(outError, err)
				outError = multierror.Append(outError, err)
			} else {
				itemResult.Hash2.Hash = h2
			}

			if itemError != nil {
				itemResult.Msg = outError.Error()
				retList = append(retList, itemResult)
				notEqualCount++
				continue
			}

			if h1 != h2 {
				//itemResult.Msg = fmt.Sprintf("hash not equal [ %v != %v ]", h1, h2)
				itemResult.Msg = "not equal hash"
				retList = append(retList, itemResult)
				notEqualCount++
				continue
			}

			//The same files
		} else {
			//The same directories
		}

		itemResult.Equal = true
		itemResult.Msg = ""
		retList = append(retList, itemResult)

		//Delete the match files
		delete(data2, k)
	}

	if len(data2) != 0 {
		for k, v2 := range data2 {
			itemResult := vo.NewCmpInfo(k, k)
			itemResult.Hash2.Size = v2.Info.Size()
			itemResult.Equal = false
			itemResult.Hash1.Error = fmt.Errorf("not exist [%v]", k)
			//itemResult.Msg = fmt.Sprintf("path1 not exist [%v]", k)
			itemResult.Msg = "not exist path1"
			retList = append(retList, itemResult)
			notEqualCount++
		}
	}

	return notEqualCount == 0, retList, nil
}

func CmpFile(file1, file2 string) (bool, vo.CmpInfo, error) {
	result := vo.NewCmpInfo(file1, file2)
	var outError error

	f1, err := os.Stat(file1)
	if err != nil {
		result.Hash1.Error = err
		outError = multierror.Append(outError, err)
	} else {
		result.Hash1.Size = f1.Size()
	}

	f2, err := os.Stat(file2)
	if err != nil {
		result.Hash2.Error = err
		outError = multierror.Append(outError, err)
	} else {
		result.Hash2.Size = f2.Size()
	}

	h1, err := util.Sha256(file1)
	if err != nil {
		result.Hash1.Error = err
		outError = multierror.Append(outError, err)
	} else {
		result.Hash1.Hash = h1
	}

	h2, err := util.Sha256(file2)
	if err != nil {
		result.Hash2.Error = err
		outError = multierror.Append(outError, err)
	} else {
		result.Hash2.Hash = h2
	}

	result.Equal = false
	if outError != nil {
		result.Msg = outError.Error()
		return false, result, outError
	}

	if f1.Size() != f2.Size() {
		//result.Msg = fmt.Sprintf("size not equal [ %v != %v ]", f1.Size(), f2.Size())
		result.Msg = "not equal size"
		return false, result, outError
	}

	if h1 != h2 {
		//result.Msg = fmt.Sprintf("hash not equal [ %v != %v ]", h1, h2)
		result.Msg = "not equal hash"
		return false, result, outError
	}

	result.Equal = true
	result.Msg = "equal"

	return true, result, nil
}
