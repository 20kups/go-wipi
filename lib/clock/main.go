package clock

import (
	"time"

	. "go-wipi/lib"
	. "go-wipi/lib/models"
)

type Clock struct {
	gpioPin int
	pulse time.Duration
	taskQueue *TaskQueue
}

func NewClock(pin int, pulseMs int, tq *TaskQueue) *Clock {
	return &Clock{
		gpioPin: pin,
		pulse: time.Duration(pulseMs) * time.Millisecond,
		taskQueue: tq,
	}
}

func(c *Clock) Init() {
	go func() {
		for {
			task, ok := c.taskQueue.Dequeue()
			if ok {
				task.Run()
				if task.Type() != TaskTypeLatchAndClear {
					DigitalWrite(c.gpioPin, DigitalValueHigh)
					time.Sleep(c.pulse/2)
					
					DigitalWrite(c.gpioPin, DigitalValueLow)
					time.Sleep(c.pulse/2)
				}
			}
		}
	}()
}
