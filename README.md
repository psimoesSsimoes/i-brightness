# MBAir-Brightness

mbair-brightness: Adjust brightness on a Macbook Air 2015.

## Why it's usefull

On a Macbook Air 2015 the keyboard brightness controls don't work automatically (as far as Fedora 27 ,anyway).
This model uses an Intel graphics card and you can manually adjust the brightness on the CLI, as root, by echoing the number into /sys/class/backlight/intel_backlight/brightness. This script is a helper that does exactly that.


## Setup

	go get /github.com/psimoesSsimoes/mbair-brightness

## Usage

	set : set brigthness manually. Accepted values between 90 and 1280

	up : increase brigthness by 10

	down : decrease brightness by 10
	

## Author

Pedro Simões (aka seomis)
