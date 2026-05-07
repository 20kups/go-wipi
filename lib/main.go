package lib

/**
 *
 * A set of golang wrapper functions around the
 * wiringPi shared library found in '/usr/lib/libwiringPi.so'
 *
**/

/*
#cgo LDFLAGS: -L/usr/lib -lwiringPi

int wiringPiSetup();
int wiringPiSetupGpio();

int softPwmCreate (int pin, int initialValue, int pwmRange);
void softPwmWrite (int pin, int value);

int digitalRead (int pin);
void digitalWrite (int pin, int value);

void pinMode (int pin, int mode);

void pullUpDnControl (int pin, int pud);

int softToneCreate (int pin);
void softToneWrite (int pin, int freq);
void softToneStop (int pin);
*/
import "C"
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type Mode int
const (
	ModeInput Mode = iota
	ModeOutput
)

type DigitalValue int
const (
	DigitalValueLow DigitalValue = iota
	DigitalValueHigh
)

type PUD int
const (
	PudOff PUD = iota
	PudDown
	PudUp
)

func WiringPiSetup() int {
	return int(C.wiringPiSetup())
}

func WiringPiSetupGpio() int {
	return int(C.wiringPiSetupGpio())
}

func WiringPiCleanup(termFunc func()) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println()
		fmt.Println("SIGINT : Shutting Down")

		termFunc()

		os.Exit(0)
	}()
}

func SoftPwmCreate(pin, initialValue, pwmRange int) int {
	return int(C.softPwmCreate(C.int(pin), C.int(initialValue), C.int(pwmRange)))
}

func SoftPwmWrite(pin, value int) {
	C.softPwmWrite(C.int(pin), C.int(value))
}

func DigitalRead(pin int) DigitalValue {
	return DigitalValue(C.digitalRead(C.int(pin)))
}
func DigitalWrite(pin int, dv DigitalValue) {
	C.digitalWrite(C.int(pin), C.int(dv))
}

func PinMode(pin int, m Mode) {
	C.pinMode(C.int(pin), C.int(m))
}

func PullUpDnControl(pin int, pud PUD) {
	C.pullUpDnControl(C.int(pin), C.int(pud));
}

func SoftToneCreate(pin int) int {
	return int(C.softToneCreate(C.int(pin)))
}
func SoftToneStop(pin int)  {
	C.softToneStop(C.int(pin))
}
func SoftToneWrite(pin, freq int) {
	C.softToneWrite(C.int(pin), C.int(freq));
}