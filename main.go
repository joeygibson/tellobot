package main

import (
	"time"

	"github.com/sirupsen/logrus"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

func main() {
	drone := tello.NewDriver("8888")

	drone.On(tello.TakeoffEvent, onTakeOff)
	drone.On(tello.LandingEvent, onLanding)

	work := func() {
		flightPlan(drone)
	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work,
	)

	robot.Start()
}

func onTakeOff(s interface{}) {
	logrus.WithField("payload", s).Info("Takeoff:")
}

func onLanding(s interface{}) {
	logrus.WithField("payload", s).Info("Landing:")
}

func flightPlan(drone *tello.Driver) {
	err := drone.StartVideo()
	if err != nil {
		logrus.WithField("err", err).Error("starting camera")
	}

	err = drone.TakeOff()
	if err != nil {
		logrus.WithField("err", err).Error("take off")
	}

	err = drone.BackFlip()
	if err != nil {
		logrus.WithField("err", err).Error("back flip")
	}

	gobot.After(5*time.Second, func() {
		err := drone.Land()
		if err != nil {
			logrus.WithField("err", err).Error("landing")
		}
	})
}
