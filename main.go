package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/stianeikeland/go-rpio"
)
var (
	exit = false
	result string

	// Buttons in faceplate
	KEY_UP     = rpio.Pin(6) 
	KEY_DOWN   = rpio.Pin(19)
	KEY_LEFT   = rpio.Pin(5)
	KEY_RIGHT  = rpio.Pin(26)
	KEY_PRESS  = rpio.Pin(13)
	KEY_BTN1   = rpio.Pin(21)
	KEY_BTN2   = rpio.Pin(20)
	KEY_BTN3   = rpio.Pin(16)
)

func main() {
	initGpio()
	for !exit {
		// Max 16 println statements with length of 16
		output := []string{}
		output = append(output, time.Now().Local().Format("01-02       1504"))
		output = append(output, "KEY_1     " + strconv.Itoa(int(KEY_BTN1.Read())))
		output = append(output, "KEY_2     " + strconv.Itoa(int(KEY_BTN2.Read())))
		output = append(output, "KEY_3     " + strconv.Itoa(int(KEY_BTN3.Read())))
		output = append(output, "KEY_CLK   " + strconv.Itoa(int(KEY_PRESS.Read())))
		output = append(output, "KEY_UP    " + strconv.Itoa(int(KEY_UP.Read())))
		output = append(output, "KEY_DOWN  " + strconv.Itoa(int(KEY_DOWN.Read())))
		output = append(output, "KEY_LEFT  " + strconv.Itoa(int(KEY_LEFT.Read())))
		output = append(output, "KEY_RIGHT " + strconv.Itoa(int(KEY_RIGHT.Read())))
		
		time.Sleep(100*time.Millisecond)
		write(output)
		//exit = true
	}
	rpio.Close()
}

func write(input []string) {
	if len(input) > 16 {
		log.Fatalln("ERR: PRINTING ARRAY LARGER THAN 16")
		os.Exit(2)
	} else if len(input) < 16 {
		input = append(input, string(len(input)))
		for len(input) != 16 {
			input = append(input, "")
		}
	}



	output := strings.Join(input, "\n")
	print("\033[2J")
	print(output)
	//print(string(len(input))+"\n")
	//fmt.Printf("%v\n", input)
}

func initGpio() {
/* 
	//Python Reference
	GPIO.setup(KEY_UP_PIN,      GPIO.IN, pull_up_down=GPIO.PUD_UP)
	GPIO.setup(KEY_DOWN_PIN,    GPIO.IN, pull_up_down=GPIO.PUD_UP)
	GPIO.setup(KEY_LEFT_PIN,    GPIO.IN, pull_up_down=GPIO.PUD_UP)
	GPIO.setup(KEY_RIGHT_PIN,   GPIO.IN, pull_up_down=GPIO.PUD_UP)
	GPIO.setup(KEY_PRESS_PIN,   GPIO.IN, pull_up_down=GPIO.PUD_UP)
	GPIO.setup(KEY1_PIN,        GPIO.IN, pull_up_down=GPIO.PUD_UP)
	GPIO.setup(KEY2_PIN,        GPIO.IN, pull_up_down=GPIO.PUD_UP)
	GPIO.setup(KEY3_PIN,        GPIO.IN, pull_up_down=GPIO.PUD_UP)
*/


	err := rpio.Open()
	if err != nil {
		print("ERROR READING GPIO")
	}
	// Initialize pins
	KEY_UP.Input()
	KEY_DOWN.Input()
	KEY_LEFT.Input()
	KEY_RIGHT.Input()
	KEY_PRESS.Input()
	KEY_BTN1.Input()
	KEY_BTN2.Input()
	KEY_BTN3.Input()

	KEY_UP.PullUp()
	KEY_DOWN.PullUp()
	KEY_LEFT.PullUp()
	KEY_RIGHT.PullUp()
	KEY_PRESS.PullUp()
	KEY_BTN1.PullUp()
	KEY_BTN2.PullUp()
	KEY_BTN3.PullUp()
}
