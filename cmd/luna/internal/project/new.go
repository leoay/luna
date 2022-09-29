package project

import (
	"bufio"
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/leoay/luna/cmd/luna/v2/internal/base"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// Project is a project template.
type Project struct {
	Name string
	Path string
}

// 获取指定目录下的所有文件,包含子目录下的文件
func GetAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetAllFiles(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".go")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}

// 修改Makefile
func MotifyMakefile(filepath string, target string) error {
	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	//读取文件内容到io中
	reader := bufio.NewReader(file)
	pos := int64(0)

	for {
		//读取每一行内容
		line, err := reader.ReadString('\n')
		if err != nil {
			//读到末尾
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		//根据关键词覆盖当前行
		if strings.Contains(line, "luna-layout") {
			fmt.Println("PPPPPP")
			k := 0
			spaceStr := ""
			for k < len(line)-1 {
				spaceStr = spaceStr + " "
				k++
			}
			file.WriteAt([]byte(spaceStr), pos)
			fmt.Println("SDSDDSDDD: ", strings.Split(line, "luna-layout")[0]+target+"\n")
			bytes := []byte(strings.Split(line, "luna-layout")[0] + target)
			file.WriteAt(bytes, pos)
		}
		//每一行读取完后记录位置
		pos += int64(len(line))
	}
	return nil
}

func (p *Project) New(ctx context.Context, dir string, layout string, branch string) error {
	to := path.Join(dir, p.Name)
	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Printf("🚫 %s already exists\n", p.Name)
		override := false
		prompt := &survey.Confirm{
			Message: "📂 Do you want to override the folder ?",
			Help:    "Delete the existing folder and create the project.",
		}
		e := survey.AskOne(prompt, &override)
		if e != nil {
			return e
		}
		if !override {
			return err
		}
		os.RemoveAll(to)
	}
	fmt.Printf("🚀 Creating service %s, layout repo is %s, please wait a moment.\n\n", p.Name, layout)
	repo := base.NewRepo(layout, branch)
	if err := repo.CopyTo(ctx, to, p.Path, []string{".git", ".github"}); err != nil {
		return err
	}
	e := os.Rename(
		path.Join(to, "cmd", "luna-layout"),
		path.Join(to, "cmd", p.Name),
	)
	if e != nil {
		return e
	}
	base.Tree(to, dir)

	files, err := GetAllFiles(to + "/internal")
	if err != nil {
		return err
	}

	files = append(files, to+"/cmd/"+p.Name+"/main.go")

	files2, err := GetAllFiles(to + "/pkg")
	if err != nil {
		return err
	}

	files = append(files, files2...)

	for _, v := range files {
		//读写方式打开文件
		file, err := os.OpenFile(v, os.O_RDWR, 0666)
		if err != nil {
			fmt.Println("open file filed.", err)
			return err
		}
		//defer关闭文件
		defer file.Close()
		//获取文件大小
		stat, err := file.Stat()
		if err != nil {
			panic(err)
		}
		var size = stat.Size()
		fmt.Println("file size:", size)

		//读取文件内容到io中
		reader := bufio.NewReader(file)
		pos := int64(0)

		for {
			//读取每一行内容
			line, err := reader.ReadString('\n')
			if err != nil {
				//读到末尾
				if err == io.EOF {
					break
				} else {
					return err
				}
			}
			//根据关键词覆盖当前行
			if strings.Contains(line, "\"luna-layout") {
				k := 0
				spaceStr := ""
				for k < len(line) {
					spaceStr = spaceStr + " "
					k++
				}
				file.WriteAt([]byte(spaceStr), pos)
				bytes := []byte(strings.Split(line, "luna-layout/")[0] + p.Name + "/" + strings.Split(line, "\"luna-layout/")[1])
				file.WriteAt(bytes, pos)
			}
			//每一行读取完后记录位置
			pos += int64(len(line))
		}
	}

	//修改makefile
	makefilePath := to + "/Makefile"
	MotifyMakefile(makefilePath, p.Name)

	fmt.Printf("\n🍺 Project creation succeeded %s\n", color.GreenString(p.Name))
	fmt.Print("💻 Use the following command to start the project 👇:\n\n")

	fmt.Println(color.WhiteString("$ cd %s", p.Name))
	fmt.Println(color.WhiteString("$ go generate ./..."))
	fmt.Println(color.WhiteString("$ go build -o ./bin/ ./... "))
	fmt.Println(color.WhiteString("$ ./bin/%s -conf ./configs\n", p.Name))
	fmt.Println("			🤝 Thanks for using Luna")
	fmt.Println("	📚 Tutorial: ")
	return nil
}
