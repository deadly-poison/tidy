package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "path"
  "path/filepath"
  "strings"
)

func getFileList(p string) {
	err := filepath.Walk(p, func(path string, f os.FileInfo, err error) error {
		if f == nil {return err}
		if f.IsDir() {return nil}
		if !strings.Contains(path, ".idea"){
			if !strings.Contains(path, "main.go"){
        strs := strings.Split(path, "/")
        newFile := filepath.Join(p, strs[len(strs)-1])
				err = os.Rename(path, newFile)
				if err != nil{
					fmt.Println(err)
				}
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func findEmptyFolder(dirname string) (emptys []string, err error) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}
	// 判断底下是否有文件
	if len(files) == 0 {
		return []string{dirname}, nil
	}

	for _, file := range files {
		if file.IsDir() {
			edirs, err := findEmptyFolder(path.Join(dirname, file.Name()))
			if err != nil {
				return nil, err
			}
			if edirs != nil {
				emptys = append(emptys, edirs...)
			}
		}
	}
	return emptys, nil
}


func main(){
  ex,err:=os.Executable()
  if err!=nil{
    panic(err)
  }
  //	获取执行文件所在目录
  exPath := filepath.Dir(ex)

	getFileList(exPath)
	emps, err := findEmptyFolder(exPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, dir := range emps {
		if err := os.Remove(dir); err != nil {
			fmt.Println("错误:", err.Error())
		} else {
			fmt.Println("删除成功:", dir)
		}
	}
}
