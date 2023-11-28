package activity

import (
	"time"
)

const (
	SportGeneric    = "Generic"
	SportHiking     = "Hiking"
	SportWalking    = "Walking"
	SportRunning    = "Running"
	SportSwimming   = "Swimming"
	SportTransition = "Transition" // Multisport transition
)

func IsConsideredMoving(sport string, speed *float64) bool {
	if speed == nil {
		return false
	}

	switch sport {
	case SportGeneric:
		fallthrough // Since we don't know the specific sport, let's assume the slowest possible speed as our tolerance.
	case SportHiking, SportWalking, SportSwimming, SportTransition: // slow moving exercise
		return *speed > 0.1388 // = 0.5km/h
	case SportRunning:
		return *speed > 0.7916 // = 2.85 km/h
	default:
		return *speed > 1.41 // 5.07 km/h
	}
}

func HasPace(sport string) bool {
	switch sport {
	case SportGeneric:
		// Since we don't know the specific sport, let's calculate the pace for now and allow the user to decide for themselves.
		// It's better to provide additional information than to have none. (TBD)
		fallthrough
	case SportHiking, SportWalking, SportRunning, SportSwimming, SportTransition:
		return true
	default:
		return false
	}
}

func isBelong(timestamp, startTime, endTime time.Time) bool {
	if timestamp.Equal(startTime) {
		return true
	}
	if endTime.IsZero() { // Last Lap or Session has no EndTime
		return timestamp.After(startTime)
	}
	if timestamp.Equal(endTime) {
		return true
	}
	return timestamp.After(startTime) && timestamp.Before(endTime)
}
