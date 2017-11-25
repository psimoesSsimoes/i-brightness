package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
)

func IsRoot() bool{
	return os.Geteuid()==0
}

func Inc(c *cli.Context)error {
	if !IsRoot(){
		return fmt.Errorf("this tool needs root access")
	}
	fmt.Println("increase")
	return nil
}
func Dec(c *cli.Context)error {
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


   
 app.Commands = []cli.Command{
    {
      Name:    "set",
      Aliases: []string{"s"},
      Usage:   "set brightness manually",
      Action:  Set,
    },
    {
      Name:    "inc",
      Aliases: []string{"i"},
      Usage:   "increase brightness",
      Action:  Inc,
  },
	 {
      Name:    "dec",
      Aliases: []string{"d"},
      Usage:   "decrease brightness",
      Action:  Dec,
    },
  }
  app.Version="0.0.1"
  app.Author="Seomis : psimoes@campus.fct.unl.pt"
  app.Run(os.Args)
}
