package jzip

import (
	"path/filepath"

	"github.com/culturadevops/jgt/exe"
)

type Jzip struct {
	RepoName  string
	FinalPath string
	Branch    string
	Die       bool
	Jexe      *exe.Jexe
}

func (g *Jzip) PrepareInit() {
	g.Jexe = new(exe.Jexe)
	g.Jexe.PrepareDefaultjExe("zip")

}
func (g *Jzip) Command(arg ...string) {
	g.Jexe.Arg = arg
	g.Jexe.CommandInternal(true)
	if g.FinalPath != "" {
		absPath, _ := filepath.Abs(g.FinalPath)
		g.Jexe.Cmd.Dir = absPath
	}
	g.Jexe.Run(g.Die)
}
func (g *Jzip) ZipFolder(NameZip string, FolderZip string) {
	g.Command("-r", NameZip, FolderZip)
}
