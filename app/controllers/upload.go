package controllers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/revel/revel"
)

type UploadSupport struct {
	*revel.Controller
}

type UploadResult struct {
	Success bool   `json:"success"`
	File    string `json:"file"`
	Error   string `json:"error"`
}

func (c UploadSupport) UploadFile(qqfile []byte) revel.Result {
	errMsg := ""
	if err := os.MkdirAll("tmp/files", os.ModePerm); err != nil {
		errMsg = err.Error()
		revel.ERROR.Print(err)
	} else {
		if len(c.Params.Files["qqfile"]) > 0 {
			filename := c.Params.Files["qqfile"][0].Filename
			if filename, err := c.ensureFileName("tmp/files", filename); err == nil {
				if writeError := ioutil.WriteFile("tmp/files/"+filename, qqfile, os.ModeExclusive); writeError != nil {
					errMsg = writeError.Error()
					revel.ERROR.Print(writeError)
				} else {
					return c.RenderJSON(UploadResult{true, filename, ""})
				}
			} else {
				errMsg = err.Error()
			}
		}
	}

	return c.RenderJSON(UploadResult{false, "", errMsg})
}

func (c UploadSupport) ensureFileName(dir string, file string) (string, error) {
	parts := strings.Split(file, ".")

	if len(parts) != 2 {
		return file, errors.New("invliad file name")
	}
	filename, ext := parts[0], parts[1]

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return dir, err
	}

	suffixes := []int64{}
	for _, fi := range files {
		if !fi.IsDir() {
			reg := regexp.MustCompile(filename + `(\(\d+\))?\.` + ext)
			matches := reg.FindStringSubmatch(fi.Name())
			if len(matches) == 2 {
				var idx int64
				if matches[1] != "" {
					idx, _ = strconv.ParseInt(matches[1][1:len(matches[1])-1], 10, 64)
				}

				suffixes = append(suffixes, idx)
			}
		}
	}

	if len(suffixes) > 0 {
		sort.Slice(suffixes, func(i, j int) bool { return suffixes[i] < suffixes[j] })
		return fmt.Sprintf("%s(%d).%s", filename, suffixes[len(suffixes)-1]+1, ext), nil
	}
	return file, nil
}
