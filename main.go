package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

var brightness = "/sys/class/backlight/intel_backlight/brightness"
var maxbrightness = "/sys/class/backlight/intel_backlight/max_brightness"

//IsRoot checks if user is running command as root
func IsRoot() bool {
	return os.Geteuid() == 0
}

//Inc function opens brightness file, reads, and increments by 50, if possible
func Inc(c *cli.Context) error {
	if !IsRoot() {
		fmt.Println("Got root?")
		return fmt.Errorf("this tool needs root access")
	}
	fmt.Println("increase")
	return nil
}

//Dec function opens brightness file, reads, and decrements by 50, if possible
func Dec(c *cli.Context) error {
	if !IsRoot() {
		fmt.Println("Got root?")
		return fmt.Errorf("this tool needs root access")
	}
	fmt.Println("decrease")
	return nil
}

//Set write to brightness file a value corresponding to the percentage given by the user
func Set(c *cli.Context) error {

	var (
		err        error
		percentage int
		max        int
		value      int
	)
	if !IsRoot() {
		fmt.Println("Got root?")
		return fmt.Errorf("this tool needs root access")
	}

	if c.NArg() > 0 {

		arg := c.Args()[0]

		percentage, err = strconv.Atoi(arg)

		if err != nil {

			fmt.Println("Percentage has to be an integer")

			return fmt.Errorf("Percentage has to be an integer")
		}

		switch {

		case percentage < 20 || percentage > 100:

			fmt.Println("Percentage has to be between 20 and 100")

			return fmt.Errorf("Percentage has to be between 20 and 100")

		default:
			max, err = ReadFile(maxbrightness)

			if err != nil {
				fmt.Println("failed to find max brightness")

				return fmt.Errorf("failed to find max brightness")

			}
			value = ((max - 500) * percentage) / 100
			fmt.Println(value)
		}
	}

	return WriteFile(value)
}

//ReadFile reads a file containing only integers
func ReadFile(filename string) (int, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error")
		return -1, err
	}
	fmt.Println(string(content[:]))
	return strconv.Atoi(strings.TrimSpace(string(content[:])))
}

//WriteFile writes an integer to brightness file
func WriteFile(n int) error {

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
	app.Version = "0.0.1"
	app.Author = "Seomis : psimoes@campus.fct.unl.pt"
	app.Run(os.Args)
}
