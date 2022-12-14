package project

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/leoay/luna/cmd/luna/v2/internal/base"
)

var repoAddIgnores = []string{
	".git", ".github", "api", "README.md", "LICENSE", "go.mod", "go.sum", "third_party", "openapi.yaml", ".gitignore",
}

func (p *Project) Add(ctx context.Context, dir string, layout string, branch string, mod string) error {
	to := path.Join(dir, p.Path)

	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Printf("๐ซ %s already exists\n", p.Name)
		override := false
		prompt := &survey.Confirm{
			Message: "๐ Do you want to override the folder ?",
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

	fmt.Printf("๐ Add service %s, layout repo is %s, please wait a moment.\n\n", p.Name, layout)

	repo := base.NewRepo(layout, branch)

	if err := repo.CopyToV2(ctx, to, path.Join(mod, p.Path), repoAddIgnores, []string{path.Join(p.Path, "api"), "api"}); err != nil {
		return err
	}

	e := os.Rename(
		path.Join(to, "cmd", "server"),
		path.Join(to, "cmd", p.Name),
	)
	if e != nil {
		return e
	}

	base.Tree(to, dir)

	//้ๅๆไปถ๏ผไฟฎๆนๅจๅฑ

	fmt.Println(color.WhiteString("$ cd %s", p.Name))
	fmt.Printf("\n๐บ Repository creation succeeded %s\n", color.GreenString(p.Name))
	fmt.Print("๐ปAAAA Use the following command to add a project ๐:\n\n")

	fmt.Println(color.WhiteString("$ cd %s", p.Name))
	fmt.Println(color.WhiteString("$ go generate ./..."))
	fmt.Println(color.WhiteString("$ go build -o ./bin/ ./... "))
	fmt.Println(color.WhiteString("$ ./bin/%s -conf ./configs\n", p.Name))
	fmt.Println("			๐ค Thanks for using Kratos")
	fmt.Println("	๐ Tutorial: ")
	return nil
}
