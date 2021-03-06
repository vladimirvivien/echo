package gexe

import (
	"regexp"

	"github.com/vladimirvivien/gexe/prog"
	"github.com/vladimirvivien/gexe/vars"
)

var (
	DefaultEcho = New()
	spaceRgx    = regexp.MustCompile("\\s")
	lineRgx     = regexp.MustCompile("\\n")
)

// Echo represents a new Echo session
type Echo struct {
	err  error
	vars *vars.Variables // session vars
	prog *prog.ProgInfo
	Conf *Config         // session config
}

// New creates a new Echo session
func New() *Echo {
	e := &Echo{
		vars: vars.New(),
		prog: prog.Prog(),
		Conf: &Config{escapeChar: '\\'},
	}
	return e
}
