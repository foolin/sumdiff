package sumdiff

import (
	"fmt"
	"github.com/foolin/sumdiff/internal/statusbar"
	"github.com/foolin/sumdiff/internal/util"
	"github.com/foolin/sumdiff/internal/vo"
	"github.com/hashicorp/go-multierror"
	"os"
)

func Cmp(path1, path2 string) (bool, []*vo.CmpVo, error) {
	result := vo.NewCmpVo(path1, path2)
	result.OK = false
	var outError *multierror.Error
	s1, err := os.Stat(path1)
	if err != nil {
		result.X.Error = err
		outError = multierror.Append(outError, err)
	}
	s2, err := os.Stat(path2)
	if err != nil {
		result.Y.Error = err
		outError = multierror.Append(outError, err)
	}
	if outError != nil {
		result.Msg = outError.Error()
		return false, []*vo.CmpVo{result}, outError
	}
	if s1.IsDir() != s2.IsDir() {
		result.Msg = fmt.Sprintf("file type not equal [ %v != %v ]", util.FileType(s1.IsDir()), util.FileType(s2.IsDir()))
		return false, []*vo.CmpVo{result}, outError
	}

	if s1.IsDir() {
		return CmpDir(path1, path2)
	} else {
		ok, ret, err := CmpFile(path1, path2)
		return ok, []*vo.CmpVo{ret}, err
	}
}

func CmpDir(path1, path2 string) (bool, []*vo.CmpVo, error) {
	var outError *multierror.Error
	outResult := vo.NewCmpVo(path1, path2)
	data1, err := util.ListPath(path1, false)
	if err != nil {
		outResult.X.Error = err
		outError = multierror.Append(outError, err)
	}
	data2, err := util.ListPath(path2, false)
	if err != nil {
		outResult.Y.Error = err
		outError = multierror.Append(outError, err)
	}
	if outError != nil {
		outResult.OK = false
		outResult.Msg = outError.Error()
		return false, []*vo.CmpVo{outResult}, outError
	}

	notEqualCount := 0
	retList := make([]*vo.CmpVo, 0)
	for k, v1 := range data1 {
		statusbar.Display("compare path " + k)
		itemResult := vo.NewCmpVo(k, k)
		itemResult.X.Size = v1.Info.Size()
		itemResult.OK = false
		v2, ok := data2[k]
		if !ok {
			itemResult.Y.Error = fmt.Errorf("not exist [%v]", k)
			itemResult.Msg = fmt.Sprintf("path2 not exist [%v]", k)
			retList = append(retList, itemResult)
			notEqualCount++
			continue
		} else {
			itemResult.Y.Size = v2.Info.Size()
		}
		if v1.Info.Size() != v2.Info.Size() {
			itemResult.Msg = fmt.Sprintf("size not equal [ %v != %v ]", v1.Info.Size(), v2.Info.Size())
			retList = append(retList, itemResult)
			notEqualCount++
			continue
		}
		if v1.Info.IsDir() != v2.Info.IsDir() {
			itemResult.Msg = fmt.Sprintf("file type not equal [ %v != %v ]", util.FileType(v1.Info.IsDir()), util.FileType(v2.Info.IsDir()))
			retList = append(retList, itemResult)
			notEqualCount++
			continue
		}
		//File check file md5
		if !v1.Info.IsDir() {
			var itemError *multierror.Error
			h1, err := util.Sha256(v1.Path)
			if err != nil {
				itemResult.X.Error = err
				itemError = multierror.Append(outError, err)
				outError = multierror.Append(outError, err)
			} else {
				itemResult.X.Hash = h1
			}
			h2, err := util.Sha256(v2.Path)
			if err != nil {
				itemResult.Y.Error = err
				itemError = multierror.Append(outError, err)
				outError = multierror.Append(outError, err)
			} else {
				itemResult.Y.Hash = h2
			}
			if itemError != nil {
				itemResult.Msg = outError.Error()
				retList = append(retList, itemResult)
				notEqualCount++
				continue
			}
			if h1 != h2 {
				itemResult.Msg = fmt.Sprintf("hash not equal [ %v != %v ]", h1, h2)
				retList = append(retList, itemResult)
				notEqualCount++
				continue
			}

			//文件相同
		} else {
			//目录相同
		}

		itemResult.OK = true
		itemResult.Msg = ""
		retList = append(retList, itemResult)

		//删除data2，后面判断是否有剩余的文件
		delete(data2, k)
	}
	if len(data2) != 0 {
		for k, v2 := range data2 {
			itemResult := vo.NewCmpVo(k, k)
			itemResult.Y.Size = v2.Info.Size()
			itemResult.OK = false
			itemResult.X.Error = fmt.Errorf("not exist [%v]", k)
			itemResult.Msg = fmt.Sprintf("path1 not exist [%v]", k)
			retList = append(retList, itemResult)
			notEqualCount++
		}
	}
	return notEqualCount == 0, retList, nil
}

func CmpFile(file1, file2 string) (bool, *vo.CmpVo, error) {
	result := vo.NewCmpVo(file1, file2)
	var outError *multierror.Error

	f1, err := os.Stat(file1)
	if err != nil {
		result.X.Error = err
		outError = multierror.Append(outError, err)
	} else {
		result.X.Size = f1.Size()
	}
	f2, err := os.Stat(file2)
	if err != nil {
		result.Y.Error = err
		outError = multierror.Append(outError, err)
	} else {
		result.Y.Size = f2.Size()
	}
	h1, err := util.Sha256(file1)
	if err != nil {
		result.X.Error = err
		outError = multierror.Append(outError, err)
	}
	h2, err := util.Sha256(file2)
	if err != nil {
		result.Y.Error = err
		outError = multierror.Append(outError, err)
	}
	result.OK = false
	if outError != nil {
		result.Msg = outError.Error()
		return false, result, outError
	}
	if f1.Size() != f2.Size() {
		result.Msg = fmt.Sprintf("size not equal [ %v != %v ]", f1.Size(), f2.Size())
		return false, result, outError
	}
	if h1 != h2 {
		result.Msg = fmt.Sprintf("hash not equal [ %v != %v ]", h1, h2)
		return false, result, outError
	}

	result.OK = true
	result.Msg = "equal"
	return true, result, nil
}
