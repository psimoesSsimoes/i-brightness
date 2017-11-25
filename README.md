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
	

## Bind it to a key

i always use tilted WM like awesome or i3. For those, use [xev](https://en.wikipedia.org/wiki/Xev)to discover the id's of the keys you want to map.

Having that, bind the keys with the generated binary.

i3 example:

	bindsym <key_id> exec "sudo ~/go/bin/./mbair-bright up"
	bindsym <key_id> exec "sudo ~/go/bin/./mbair-bright down"


if your configuration doesn't execute sudo without asking the password, then you can use [capabilities](http://man7.org/linux/man-pages/man7/capabilities.7.html) to give permissions to mbair-bright binary.

## Author

Pedro Simões (aka seomis)
