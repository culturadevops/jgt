package jgit

import (
	"inkafarma/devops_tools/exe"
	"os/exec"
	"path/filepath"
)

type Jgit struct {
	RepoName  string
	FinalPath string
	Branch    string
}

func (g *Jgit) Command(arg ...string) {

	cmd := exec.Command("git", arg...)
	if g.FinalPath != "" {
		absPath, _ := filepath.Abs(g.FinalPath)
		cmd.Dir = absPath
		println(absPath)
	}
	exe.Run(cmd, true)
}
func (g *Jgit) CloneB(RepoName string, FinalPath string) {
	if FinalPath != "" {
		g.Command("clone", "-b", g.Branch, RepoName, FinalPath)
		g.FinalPath = FinalPath
	} else {
		g.Command("clone", RepoName)
		g.FinalPath = FinalPath
	}
}
func (g *Jgit) Clone(RepoName string, FinalPath string) {
	if FinalPath != "" {
		g.Command("clone", RepoName, FinalPath)
		g.FinalPath = FinalPath
	} else {
		g.Command("clone", RepoName)
		g.FinalPath = FinalPath
	}
}
func (g *Jgit) CloneBwithSSH(RepoName string, FinalPath string) {
	g.CloneB("git@github.com:/"+RepoName, FinalPath)
}
func (g *Jgit) CloneSSH(RepoName string, FinalPath string) {
	g.Clone("git@github.com:/"+RepoName, FinalPath)
}
func (g *Jgit) Add(path string) {
	g.Command("add", path)
}
func (g *Jgit) AddAll() {
	g.Add(".")
}
func (g *Jgit) Commit(comentario string) {
	g.Command("commit", "-m", comentario)
}
func (g *Jgit) Push(branch string) {
	g.Command("push", "origin", branch)
}
func (g *Jgit) PushAll() {

	g.Command("push", "origin", g.Branch)
}
func (g *Jgit) Checkout(branch string) {
	if branch == "" {
		g.Command("checkout", g.Branch)
	} else {
		g.Command("checkout", branch)
	}

}
func (g *Jgit) CheckoutB(branch string) {
	if branch == "" {
		g.Command("checkout", "-b", g.Branch)
	} else {
		g.Command("checkout", "-b", branch)
	}

}
