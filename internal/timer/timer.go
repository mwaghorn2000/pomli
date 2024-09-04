package timer

import (
	"errors"
	"log"
	"os/exec"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/sevlyar/go-daemon"
)

type NotificationType int

type Timer struct {
	duration int
}

func CreateTimer(duration int) (*Timer, error) {
	err := validateTime(duration)
	if err != nil {
		return nil, err
	}

	return &Timer{
		duration: duration,
	}, nil
}

func ExecuteTimer(timer Timer) {
	d := timer.duration

	cntxt := &daemon.Context{
		PidFileName: "pomli.pid",
		PidFilePerm: 0644,
		LogFileName: "pomli.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
	}

	x, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if x != nil {
		return
	}
	defer cntxt.Release()

	startTimer(time.Duration(d))
}

func startTimer(duration time.Duration) error {
	time.Sleep(duration)

	err := beeep.Notify("Pomli", "Your done working... FOR NOW!", "")
	if err != nil {
		return err
	}

	err = playSound()
	if err != nil {
		return err
	}

	return nil
}

func playSound() error {
	cmd := exec.Command("afplay", "/System/Library/Sounds/Glass.aiff")
	return cmd.Run()
}

func validateTime(time int) error {
	if time <= 0 {
		return errors.New("Timer must be set greater than 0")
	}

	return nil
}
