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
  app.Run(os.Args)
}
