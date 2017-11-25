package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
)

func IsRoot() bool{
	return os.Geteuid()==0
}

func Increase(c *cli.Context)error {
	if !IsRoot(){
		return fmt.Errorf("this tool needs root access")
	}
	fmt.Println("increase")
	return nil
}
func Decrease(c *cli.Context)error {
	if !IsRoot(){
		return fmt.Errorf("this tool needs root access")
	}
	fmt.Println("decrease")
	return nil	
}
func Set(c *cli.Context)error {
	if c.NArg() > 0 {
		name := c.Args()[0:]
	fmt.Println(name)
    }

	return nil
}
func main() {

  app := cli.NewApp()
  app.Name = "mac_air_brightness"
  app.Usage = "change screen brightness of your mac book air!"

    cli.AppHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}
USAGE:
   {{.HelpName}} command [arguments...]{{end}}
   {{if len .Authors}}
VERSION:
   {{.Version}}
   {{end}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}{{if .Copyright }}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}

`
   
 app.Commands = []cli.Command{
    {
      Name:    "set",
      Aliases: []string{"s"},
      Usage:   "set brightness manually",
      Action:  func(c *cli.Context) error {
        return nil
      },
    },
    {
      Name:    "inc",
      Aliases: []string{"i"},
      Usage:   "increase brightness",
      Action:  func(c *cli.Context) error {
        return nil
      },
  },
	 {
      Name:    "dec",
      Aliases: []string{"d"},
      Usage:   "decrease brightness",
      Action:  func(c *cli.Context) error {
        return nil
      },
    },
  }
  app.Version="0.0.1"
  app.Author="Seomis ==> psimoes@campus.fct.unl.pt"
  app.Run(os.Args)
}
