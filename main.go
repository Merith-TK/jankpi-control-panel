package main

import (
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
		// Max 16 println statements
		println(time.Now().Local().Format("01-02       1504"))
		println("KEY_1     ", KEY_BTN1.Read())
		println("KEY_2     ", KEY_BTN2.Read())
		println("KEY_3     ", KEY_BTN3.Read())
		println("KEY_CLK   ", KEY_PRESS.Read())
		println("KEY_UP	  ", KEY_UP.Read())
		println("KEY_DOWN  ", KEY_DOWN.Read())
		println("KEY_LEFT  ", KEY_LEFT.Read())
		println("KEY_RIGHT ", KEY_RIGHT.Read())

		// Result is max 16x16 characters
		clearDisplay()
		time.Sleep(200*time.Millisecond)
	}
	rpio.Close()
}

func clearDisplay() {
	// Modify this for the exact amount of \n needed
	//				1 2 3 4 5 6 7 8 9 0 2 3 4 5 6 7 
	cleanScreen := "\n\n\n\n\n\n"
	print(cleanScreen)
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
