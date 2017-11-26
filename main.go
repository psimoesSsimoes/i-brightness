package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

var brightness = "/sys/class/backlight/acpi_video0/brightness"
var maxbrightness = "/sys/class/backlight/acpi_video0/max_brightness"
var (
	err        error
	percentage int
	max        int
	value      int
	br         int
)

//IsRoot checks if user is running command as root
func IsRoot() bool {
	return os.Geteuid() == 0
}

//AdjustBrightness limited to values between 40 and 1400
func AdjustBrightness(origin, modification int) int {
	switch {

	case origin+modification < 1:

		return 1

	case origin+modification > 15:

		return 15

	default:

		return origin + modification
	}

}

//Inc function opens brightness file, reads, and increments by 50, if possible
func Inc(c *cli.Context) error {
	br, err = ReadFile(brightness)

	if err != nil {
		return err
	}

	return WriteFile(AdjustBrightness(br, 1))
}

//Dec function opens brightness file, reads, and decrements by 50, if possible
func Dec(c *cli.Context) error {
	br, err = ReadFile(brightness)

	if err != nil {
		return err
	}
	return WriteFile(AdjustBrightness(br, -1))
}

//Set write to brightness file a value corresponding to the percentage given by the user
func Set(c *cli.Context) error {

	if c.NArg() > 0 {

		arg := c.Args()[0]

		percentage, err = strconv.Atoi(arg)

		if err != nil {

			fmt.Println("Percentage has to be an integer")

			return fmt.Errorf("Percentage has to be an integer")
		}

		switch {

		case percentage < 10 || percentage > 100:

			fmt.Println("Percentage has to be between 20 and 100")

			return fmt.Errorf("Percentage has to be between 20 and 100")

		default:
			max, err = ReadFile(maxbrightness)

			if err != nil {
				fmt.Println("failed to find max brightness")

				return fmt.Errorf("failed to find max brightness")

			}
			value = ((max - 1) * percentage) / 100
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
	return strconv.Atoi(strings.TrimSpace(string(content)))
}

//WriteFile writes an integer to brightness file
func WriteFile(number int) error {
	f, err := os.OpenFile(brightness, os.O_WRONLY|os.O_TRUNC, 0444)

	if err != nil {
		return err
	}
	n, err := f.Write([]byte(strconv.Itoa(number)))
	if err == nil && n < len(([]byte(string(number)))) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err

}

func main() {

	if !IsRoot() {
		fmt.Println("Got root?")
		os.Exit(1)
	}

	app := cli.NewApp()
	app.Name = "i-brightness"
	app.Usage = "change screen brightness!"

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
