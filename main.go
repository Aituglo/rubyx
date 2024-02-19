package main

import (
	"bufio"
	"fmt"
	"github.com/aituglo/rubyx/pkg"
	"github.com/aituglo/rubyx/rubyx"
	"github.com/aituglo/rubyx/tool"
	"github.com/alecthomas/kong"
	"log"
	"os"
)

type Context struct {
	Config    *pkg.Config
	Cli       CLI
	ProgramID int
	InputData []string
}

type CLI struct {
	Programs     ProgramsCommand   `cmd:"programs" help:"Manage programs"`
	Scope        struct{}          `cmd:"scope" help:"Show scope"`
	New          ProgramCommand    `cmd:"new" help:"Create new program"`
	Remove       ProgramCommand    `cmd:"rm" help:"Remove program"`
	Set          ProgramCommand    `cmd:"set" help:"Set active program"`
	Unset        struct{}          `cmd:"unset" help:"Unset active program"`
	Current      struct{}          `cmd:"current" help:"Show current program"`
	Program      string            `name:"p" help:"Program name"`
	Subdomains   SubdomainsCommand `cmd:"subdomains" help:"Manage subdomains"`
	Ips          IpsCommand        `cmd:"ips" help:"Manage ips"`
	Technologies struct{}          `cmd:"technologies" help:"Show technologies"`
	Tool         ToolCommand       `cmd:"tool" help:"Run tool"`
	All          bool              `name:"all" help:"Get all"`
}

type SubdomainsCommand struct {
	Add        struct{} `cmd:"add" help:"Add subdomain" optional:""`
	Show       struct{} `cmd:"show" help:"Show subdomains" optional:""`
	Technology string   `name:"t" help:"Technology name"`
}

type IpsCommand struct {
	Add  struct{} `cmd:"add" help:"Add ip" optional:""`
	Show struct{} `cmd:"show" help:"Show ips" optional:""`
}

type ToolCommand struct {
	Name  string   `help:"Tool name"`
	Input []string `arg:"" type:"existingfile"`
}

type ProgramCommand struct {
	Program string `arg:"program" help:"Program name"`
}

type ProgramsCommand struct {
	Reload bool `help:"Reload programs" optional:""`
}

var (
	cli CLI
	ctx = kong.Parse(&cli)
)

func main() {
	var programId int
	config := &pkg.Config{}
	pkg.ReadConfig(config)

	if cli.Program != "" {
		programId = rubyx.GetProgramID(*config, cli.Program)
	} else if config.ActiveProgram != "" {
		programId = rubyx.GetProgramID(*config, config.ActiveProgram)
	}
	context := &Context{Config: config, ProgramID: programId, Cli: cli}

	switch ctx.Command() {
	case "programs":
		if cli.Programs.Reload {
			rubyx.ReloadPrograms(*config)
		}
		context.showPrograms()
	case "scope":
		context.showScope()
	case "new":
		rubyx.NewProgram(*config, cli.New.Program)
	case "rm":
		rubyx.DeleteProgram(*config, cli.Remove.Program)
	case "set":
		rubyx.SetProgram(*config, cli.Set.Program)
	case "unset":
		config.ActiveProgram = ""
		pkg.WriteConfig(*config)
	case "current":
		fmt.Println(config.ActiveProgram)
	case "subdomains show":
		context.showSubdomains()
	case "subdomains add":
		rubyx.AddSubdomain(*config, cli.Program, programId)
	case "technologies":
		context.showTechnologies()
	case "ips show":
		context.showIPs(cli.All)
	case "ips add":
		rubyx.AddIp(*config, cli.Program, programId)
	case "tool <input>":
		context.runTool(cli.Tool.Name)
	}
}

func (c *Context) showPrograms() {
	programs := rubyx.GetAllPrograms(*c.Config)
	for _, program := range programs {
		fmt.Println(program.Slug)
	}
}
func (c *Context) showScope() {
	scope := rubyx.GetScope(*c.Config, c.ProgramID)
	for _, domain := range scope {
		fmt.Println(domain.Scope)
	}
}
func (c *Context) showSubdomains() {
	var subdomains []pkg.Subdomain

	if c.Cli.Subdomains.Technology != "" {
		subdomains = rubyx.GetSubdomainsByTechnology(*c.Config, c.Cli.Subdomains.Technology)
	} else if c.Cli.All {
		subdomains = rubyx.GetAllSubdomains(*c.Config)
	} else {
		subdomains = rubyx.GetSubdomainsByProgram(*c.Config, c.ProgramID)
	}

	for _, domain := range subdomains {
		fmt.Println(domain.Subdomain)
	}
}

func (c *Context) showTechnologies() {
	technologies := rubyx.GetAllTechnologies(*c.Config)
	for _, technology := range technologies {
		fmt.Println(technology.Name)
	}
}

func (c *Context) showIPs(all bool) {
	var ips []pkg.Ip
	if all {
		ips = rubyx.GetAllIps(*c.Config)
	} else if c.Config.ActiveProgram != "" {
		programId := rubyx.GetProgramID(*c.Config, c.Config.ActiveProgram)
		ips = rubyx.GetIpsByProgram(*c.Config, programId)
	}
	if len(ips) > 0 {
		for _, ip := range ips {
			fmt.Println(ip.Ip)
		}
	}
}

func (c *Context) runTool(name string) {
	var data *os.File
	var err error

	for _, file := range c.Cli.Tool.Input {
		if file == "-" {
			data = os.Stdin
		} else {
			data, err = os.Open(file)
			if err != nil {
				fmt.Println(err)
			}
		}
		scanner := bufio.NewScanner(data)
		for scanner.Scan() {
			c.InputData = append(c.InputData, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	}

	if name == "wappago" {
		tool.WappaGo(*c.Config, c.InputData, c.ProgramID)
	}
}
