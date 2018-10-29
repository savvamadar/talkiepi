# talkiepi

talkiepi is a headless capable Mumble client written in Go, written for walkie talkie style use on the Pi using GPIO pins for push to talk and LED status.  This fork is made to add 2 new flags voiceon and usegpio. By default this fork doesn't initialize the Gpio.


## Sample run command

./talkiepi -server talkiepi.projectable.me:64738 -username TPI -insecure -voiceon


## 3D printable enclosure

In the stl directory are the stl files for the enclosure I have designed specifically for the Raspberry Pi B+ board layout (I am using a Raspberry Pi 3 Model B) and the PCB and components from the [US Robotics USB Speakerphone](https://www.amazon.com/USRobotics-USB-Internet-Speakerphone-USR9610/dp/B000E6IL10/ref=sr_1_1?ie=UTF8&qid=1472691020&sr=8-1&keywords=us+robotics+speakerphone).
I will be posting a blog post shortly with a full component list and build guide.  For more information regarding building a full talkiepi device, go check out my blog at [projectable.me](http://projectable.me).


## Installing talkiepi

I have put together an install guide [here](doc/README.md).


## GPIO

You can edit your pin assignments in `talkiepi.go`
```go
const (
	OnlineLEDPin       uint = 18
	ParticipantsLEDPin uint = 23
	TransmitLEDPin     uint = 24
	ButtonPin          uint = 25
)
```

Here is a basic schematic of how I am currently controlling the LEDs and pushbutton:

![schematic](doc/gpio_diagram.png "GPIO Diagram")


## Pi Zero Fixes
I have compiled libopenal without ARM NEON support so that it works on the Pi Zero. The packages can be found in the [workarounds](/workarounds/). directory of this repo, install the libopenal1 package over your existing libopenal install.


## License

MPL 2.0

## Author

- fork - [Savva Madar](https://github.com/savvamadar/talkiepi)
- talkiepi - [Daniel Chote](https://github.com/dchote)
- Barnard,Gumble Author - Tim Cooper (<tim.cooper@layeh.com>)

