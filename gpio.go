package talkiepi

import (
	"fmt"
	"time"

	"github.com/dchote/gpio"
	"github.com/stianeikeland/go-rpio"
)

func (b *Talkiepi) initGPIO() {
	// we need to pull in rpio to pullup our button pin
	b.GPIOEnabled = b.UseGpio
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		fmt.Printf("Unable to turn on GPIO\n")
		b.GPIOEnabled = false
		//return
	}

	if b.VoiceOn {
		fmt.Printf("Start transmit\n")
		b.TransmitStart()
	}

	if b.GPIOEnabled == false {
		fmt.Printf("Disabling GPIO\n")
		return
	}

	fmt.Printf("Enabling GPIO\n")
	ButtonPinPullUp := rpio.Pin(ButtonPin)
	ButtonPinPullUp.PullUp()

	rpio.Close()

	if b.VoiceOn == false {
		// unfortunately the gpio watcher stuff doesnt work for me in this context, so we have to poll the button instead
		b.Button = gpio.NewInput(ButtonPin)
		go func() {
			for {
				currentState, err := b.Button.Read()

				if currentState != b.ButtonState && err == nil {
					b.ButtonState = currentState

					if b.Stream != nil {
						if b.ButtonState == 1 {
							fmt.Printf("Button is released\n")
							b.TransmitStop()
						} else {
							fmt.Printf("Button is pressed\n")
							b.TransmitStart()
						}
					}

				}

				time.Sleep(500 * time.Millisecond)
			}
		}()
	}

	// then we can do our gpio stuff
	b.OnlineLED = gpio.NewOutput(OnlineLEDPin, false)
	b.ParticipantsLED = gpio.NewOutput(ParticipantsLEDPin, false)
	b.TransmitLED = gpio.NewOutput(TransmitLEDPin, false)
}

func (b *Talkiepi) LEDOn(LED gpio.Pin) {
	if b.GPIOEnabled == false {
		return
	}

	LED.High()
}

func (b *Talkiepi) LEDOff(LED gpio.Pin) {
	if b.GPIOEnabled == false {
		return
	}

	LED.Low()
}

func (b *Talkiepi) LEDOffAll() {
	if b.GPIOEnabled == false {
		return
	}

	b.LEDOff(b.OnlineLED)
	b.LEDOff(b.ParticipantsLED)
	b.LEDOff(b.TransmitLED)
}
