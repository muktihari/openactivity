package activity

import (
	"time"

	"github.com/muktihari/openactivity-fit/accumulator"
	"github.com/muktihari/openactivity-fit/kit"
)

type Session struct {
	Timestamp        time.Time
	StartTime        time.Time
	EndTime          time.Time
	Sport            string
	TotalMovingTime  float64
	TotalElapsedTime float64
	TotalDistance    float64
	TotalAscent      uint16
	TotalDescent     uint16
	TotalCalories    uint16
	AvgSpeed         *float64
	MaxSpeed         *float64
	AvgHeartRate     *uint8
	MaxHeartRate     *uint8
	AvgCadence       *uint8
	MaxCadence       *uint8
	AvgPower         *uint16
	MaxPower         *uint16
	AvgTemperature   *int8
	MaxTemperature   *int8
	AvgAltitude      *float64
	MaxAltitude      *float64
	AvgPace          *float64
	AvgElapsedPace   *float64

	Laps    []*Lap
	Records []*Record
}

func NewSessionFromLaps(laps []*Lap, sport string) *Session {
	if len(laps) == 0 {
		return nil
	}

	ses := &Session{
		Timestamp: laps[0].Timestamp,
		StartTime: laps[0].StartTime,
		EndTime:   laps[len(laps)-1].EndTime,
		Sport:     sport,
	}

	var (
		totalElapsedTimeAccumu = new(accumulator.Accumulator[float64])
		totalMovingTimeAccumu  = new(accumulator.Accumulator[float64])
		totalDistanceAccumu    = new(accumulator.Accumulator[float64])
		totalAscentAccumu      = new(accumulator.Accumulator[uint16])
		totalDescentAccumu     = new(accumulator.Accumulator[uint16])
		speedAvgAccumu         = new(accumulator.Accumulator[float64])
		speedMaxAccumu         = new(accumulator.Accumulator[float64])
		altitudeAvgAccumu      = new(accumulator.Accumulator[float64])
		altitudeMaxAccumu      = new(accumulator.Accumulator[float64])
		cadenceAvgAccumu       = new(accumulator.Accumulator[uint8])
		cadenceMaxAccumu       = new(accumulator.Accumulator[uint8])
		heartRateAvgAccumu     = new(accumulator.Accumulator[uint8])
		heartRateMaxAccumu     = new(accumulator.Accumulator[uint8])
		powerAvgAccumu         = new(accumulator.Accumulator[uint16])
		powerMaxAccumu         = new(accumulator.Accumulator[uint16])
		temperatureAvgAccumu   = new(accumulator.Accumulator[int8])
		temperatureMaxAccumu   = new(accumulator.Accumulator[int8])
	)

	for i := range laps {
		lap := laps[i]

		totalElapsedTimeAccumu.Collect(&lap.TotalElapsedTime)
		totalMovingTimeAccumu.Collect(&lap.TotalMovingTime)
		totalDistanceAccumu.Collect(&lap.TotalDistance)
		totalAscentAccumu.Collect(&lap.TotalAscent)
		totalDescentAccumu.Collect(&lap.TotalDescent)
		speedAvgAccumu.Collect(lap.AvgSpeed)
		speedMaxAccumu.Collect(lap.MaxSpeed)
		altitudeAvgAccumu.Collect(lap.AvgAltitude)
		altitudeMaxAccumu.Collect(lap.MaxAltitude)
		cadenceAvgAccumu.Collect(lap.AvgCadence)
		cadenceMaxAccumu.Collect(lap.MaxCadence)
		heartRateAvgAccumu.Collect(lap.AvgHeartRate)
		heartRateMaxAccumu.Collect(lap.MaxHeartRate)
		powerAvgAccumu.Collect(lap.AvgPower)
		powerMaxAccumu.Collect(lap.MaxPower)
		temperatureAvgAccumu.Collect(lap.AvgTemperature)
		temperatureMaxAccumu.Collect(lap.MaxTemperature)
	}

	if value := totalElapsedTimeAccumu.Sum(); value != nil {
		ses.TotalElapsedTime = *value
	}
	if value := totalMovingTimeAccumu.Sum(); value != nil {
		ses.TotalMovingTime = *value
	}
	if value := totalDistanceAccumu.Max(); value != nil {
		ses.TotalDistance = *value
	}
	if value := totalAscentAccumu.Sum(); value != nil {
		ses.TotalAscent = *value
	}
	if value := totalDescentAccumu.Sum(); value != nil {
		ses.TotalDescent = *value
	}

	ses.AvgSpeed = speedAvgAccumu.Avg()
	ses.MaxSpeed = speedMaxAccumu.Max()
	ses.AvgAltitude = altitudeAvgAccumu.Avg()
	ses.MaxAltitude = altitudeMaxAccumu.Max()
	ses.AvgCadence = cadenceAvgAccumu.Avg()
	ses.MaxCadence = cadenceMaxAccumu.Max()
	ses.AvgHeartRate = heartRateAvgAccumu.Avg()
	ses.MaxHeartRate = heartRateMaxAccumu.Max()
	ses.AvgPower = powerAvgAccumu.Avg()
	ses.MaxPower = powerMaxAccumu.Max()
	ses.AvgTemperature = temperatureAvgAccumu.Avg()
	ses.MaxTemperature = temperatureMaxAccumu.Max()

	if HasPace(sport) {
		var (
			paceAvgAccumu        = new(accumulator.Accumulator[float64])
			paceAvgElapsedAccumu = new(accumulator.Accumulator[float64])
		)
		for i := range laps {
			lap := laps[i]

			paceAvgAccumu.Collect(lap.AvgPace)
			paceAvgElapsedAccumu.Collect(lap.AvgElapsedPace)
		}
		ses.AvgPace = paceAvgAccumu.Avg()
		ses.AvgElapsedPace = paceAvgElapsedAccumu.Avg()
	}

	return ses
}

func (s *Session) ToMap() map[string]any {
	m := map[string]any{}

	if !s.Timestamp.IsZero() {
		m["timestamp"] = s.Timestamp.Format(time.RFC3339)
	}
	if !s.StartTime.IsZero() {
		m["startTime"] = s.StartTime.Format(time.RFC3339)
	}
	if !s.EndTime.IsZero() {
		m["endTime"] = s.EndTime.Format(time.RFC3339)
	}
	if s.Sport != "" {
		m["sport"] = s.Sport
	}

	m["totalMovingTime"] = s.TotalMovingTime
	m["totalElapsedTime"] = s.TotalElapsedTime
	m["totalDistance"] = s.TotalDistance
	m["totalAscent"] = s.TotalAscent
	m["totalDescent"] = s.TotalDescent
	m["totalCalories"] = s.TotalCalories

	if s.AvgSpeed != nil {
		m["avgSpeed"] = *s.AvgSpeed
	}
	if s.MaxSpeed != nil {
		m["maxSpeed"] = *s.MaxSpeed
	}
	if s.AvgHeartRate != nil {
		m["avgHeartRate"] = *s.AvgHeartRate
	}
	if s.MaxHeartRate != nil {
		m["maxHeartRate"] = *s.MaxHeartRate
	}
	if s.AvgCadence != nil {
		m["avgCadence"] = *s.AvgCadence
	}
	if s.MaxCadence != nil {
		m["maxCadence"] = *s.MaxCadence
	}
	if s.AvgPower != nil {
		m["avgPower"] = *s.AvgPower
	}
	if s.MaxPower != nil {
		m["maxPower"] = *s.MaxPower
	}
	if s.AvgTemperature != nil {
		m["avgTemperature"] = *s.AvgTemperature
	}
	if s.MaxTemperature != nil {
		m["maxTemperature"] = *s.MaxTemperature
	}
	if s.AvgAltitude != nil {
		m["avgAltitude"] = *s.AvgAltitude
	}
	if s.MaxAltitude != nil {
		m["maxAltitude"] = *s.MaxAltitude
	}
	if s.AvgPace != nil {
		m["avgPace"] = *s.AvgPace
	}
	if s.AvgElapsedPace != nil {
		m["avgElapsedPace"] = *s.AvgElapsedPace
	}
	if len(s.Laps) != 0 {
		laps := make([]any, len(s.Laps))
		for i := range s.Laps {
			laps[i] = s.Laps[i].ToMap()
		}
		m["laps"] = laps
	}
	if len(s.Records) != 0 {
		records := make([]any, len(s.Records))
		for i := range s.Records {
			records[i] = s.Records[i].ToMap()
		}
		m["records"] = records
	}

	return m
}

func (s *Session) IsBelongToThisSession(t time.Time) bool {
	return isBelong(t, s.StartTime, s.EndTime)
}

// PutLaps puts given laps into session and return any remaining laps that doesn't belong to this session.
func (s *Session) PutLaps(laps ...*Lap) (remainings []*Lap) {
	remainings = make([]*Lap, 0, len(laps))
	for j := range laps {
		lap := laps[j]
		if s.IsBelongToThisSession(lap.StartTime) {
			s.Laps = append(s.Laps, lap)
		} else {
			remainings = append(remainings, lap)
		}
	}
	return remainings
}

// PutRecords puts given records into session and return any remaining records that doesn't belong to this session.
func (s *Session) PutRecords(records ...*Record) (remainings []*Record) {
	remainings = make([]*Record, 0, len(records))
	for j := range records {
		rec := records[j]
		if s.IsBelongToThisSession(rec.Timestamp) {
			s.Records = append(s.Records, rec)
		} else {
			remainings = append(remainings, rec)
		}
	}
	return remainings
}

// CombineSession combines sesssion's values into targetSession.
// Every zero value in targetSession will be replaced with the corresponding value in session.
func CombineSession(targetSession, session *Session) {
	if targetSession == nil || session == nil {
		return
	}

	if targetSession.EndTime.IsZero() {
		targetSession.EndTime = session.EndTime
	}

	targetSession.TotalElapsedTime = kit.PickNonZeroValue(targetSession.TotalElapsedTime, session.TotalElapsedTime)
	targetSession.TotalMovingTime = kit.PickNonZeroValue(targetSession.TotalMovingTime, session.TotalMovingTime)
	targetSession.TotalDistance = kit.PickNonZeroValue(targetSession.TotalDistance, session.TotalDistance)
	targetSession.TotalCalories = kit.PickNonZeroValue(targetSession.TotalCalories, session.TotalCalories)
	targetSession.TotalAscent = kit.PickNonZeroValue(targetSession.TotalAscent, session.TotalAscent)
	targetSession.TotalDescent = kit.PickNonZeroValue(targetSession.TotalDescent, session.TotalDescent)
	targetSession.AvgSpeed = kit.PickNonZeroValuePtr(targetSession.AvgSpeed, session.AvgSpeed)
	targetSession.MaxSpeed = kit.PickNonZeroValuePtr(targetSession.MaxSpeed, session.MaxSpeed)
	targetSession.AvgHeartRate = kit.PickNonZeroValuePtr(targetSession.AvgHeartRate, session.AvgHeartRate)
	targetSession.MaxHeartRate = kit.PickNonZeroValuePtr(targetSession.MaxHeartRate, session.MaxHeartRate)
	targetSession.AvgCadence = kit.PickNonZeroValuePtr(targetSession.AvgCadence, session.AvgCadence)
	targetSession.MaxCadence = kit.PickNonZeroValuePtr(targetSession.MaxCadence, session.MaxCadence)
	targetSession.AvgPower = kit.PickNonZeroValuePtr(targetSession.AvgPower, session.AvgPower)
	targetSession.MaxPower = kit.PickNonZeroValuePtr(targetSession.MaxPower, session.MaxPower)
	targetSession.AvgTemperature = kit.PickNonZeroValuePtr(targetSession.AvgTemperature, session.AvgTemperature)
	targetSession.MaxTemperature = kit.PickNonZeroValuePtr(targetSession.MaxTemperature, session.MaxTemperature)
	targetSession.AvgAltitude = kit.PickNonZeroValuePtr(targetSession.AvgAltitude, session.AvgAltitude)
	targetSession.MaxAltitude = kit.PickNonZeroValuePtr(targetSession.MaxAltitude, session.MaxAltitude)
	targetSession.AvgPace = kit.PickNonZeroValuePtr(targetSession.AvgPace, session.AvgPace)
	targetSession.AvgElapsedPace = kit.PickNonZeroValuePtr(targetSession.AvgElapsedPace, session.AvgElapsedPace)
}
