package jgit

import (
	"github.com/culturadevops/jgt/exe"
	"github.com/culturadevops/jgt/jlog"
)

type Jgit struct {
	RepoName string
	Branch   string
	Jexe     *exe.Jexe
}

func (g *Jgit) PrepareInit() {
	g.Jexe = new(exe.Jexe)
	g.Jexe.PrepareDefaultjExe("git")
}
func (i *Jgit) ConfigureInitLog(IsDebug bool, PrinterLogs bool, PrinterScreen bool) {
	i.Jexe.Log = jlog.PrepareLog(IsDebug, PrinterLogs, PrinterScreen)
}
func (g *Jgit) PrepareInSilence() {
	g.Jexe = new(exe.Jexe)
	g.Jexe.PrepareDefaultWithLogSilence("git")
}
func (g *Jgit) BranchAndCloneB(Branch, RepoName, FinalPath string) {
	g.Branch = Branch
	g.CloneB(RepoName, FinalPath)
}

func (g *Jgit) CloneB(RepoName string, FinalPath string) {
	if FinalPath != "" {
		g.Jexe.ExecuteWithArg("clone", "-b", g.Branch, RepoName, FinalPath)
	} else {
		g.Jexe.ExecuteWithArg("clone", "-b", g.Branch, RepoName)
	}
}
func (g *Jgit) Clone(RepoName string, FinalPath string) {
	if FinalPath != "" {
		g.Jexe.ExecuteWithArg("clone", RepoName, FinalPath)
	} else {
		g.Jexe.ExecuteWithArg("clone", RepoName)
	}
}
func (g *Jgit) CloneBwithSSH(RepoName string, FinalPath string) {
	g.CloneB("git@github.com:/"+RepoName, FinalPath)
}
func (g *Jgit) CloneSSH(RepoName string, FinalPath string) {
	g.Clone("git@github.com:/"+RepoName, FinalPath)
}
func (g *Jgit) Add(path string) {
	g.Jexe.ExecuteWithArg("add", path)
}
func (g *Jgit) AddAll() {
	g.Add(".")
}
func (g *Jgit) Commit(comentario string) {
	g.Jexe.ExecuteWithArg("commit", "-m", comentario)
}
func (g *Jgit) Push(branch string) {
	g.Jexe.ExecuteWithArg("push", "origin", branch)
}
func (g *Jgit) PushAll() {

	g.Jexe.ExecuteWithArg("push", "origin", g.Branch)
}
func (g *Jgit) Checkout(branch string) {
	if branch == "" {
		g.Jexe.ExecuteWithArg("checkout", g.Branch)
	} else {
		g.Jexe.ExecuteWithArg("checkout", branch)
	}

}
func (g *Jgit) CheckoutB(branch string) {
	if branch == "" {
		g.Jexe.ExecuteWithArg("checkout", "-b", g.Branch)
	} else {
		g.Jexe.ExecuteWithArg("checkout", "-b", branch)
	}

}
