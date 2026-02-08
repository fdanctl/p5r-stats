package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Stat int8

const (
	Knowledge Stat = iota
	Guts
	Proficiency
	Kindness
	Charm
)

func (s Stat) String() string {
	switch s {
	case Knowledge:
		return "knowledge"
	case Guts:
		return "guts"
	case Proficiency:
		return "proficiency"
	case Kindness:
		return "kindness"
	case Charm:
		return "charm"
	default:
		return "unknown"
	}
}

func ParseStat(s string) (Stat, error) {
	switch s {
	case "knowledge":
		return Knowledge, nil
	case "guts":
		return Guts, nil
	case "proficiency":
		return Proficiency, nil
	case "kindness":
		return Kindness, nil
	case "charm":
		return Charm, nil
	default:
		return -1, fmt.Errorf("%w: %s", ErrInvalidStat, s)
	}
}

func (s *Stat) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	stat, _ := ParseStat(str)

	*s = stat
	return nil
}

func (s Stat) MarshalJSON() ([]byte, error) {
	str := s.String()
	return json.Marshal(str)
}

type IncreasedStat struct {
	Stat   Stat  `json:"stat"`
	Points uint8 `json:"points"`
}

func (v IncreasedStat) Validate() error {
	if v.Stat.String() == "unknown" {
		return ErrInvalidStat
	}
	if v.Points == 0 {
		return errors.New("Field 'points' is required, and must not be 0.")
	}
	return nil
}

type Activity struct {
	Id             string          `json:"id"`
	Title          string          `json:"title"`
	Description    string          `json:"description"`
	Date           time.Time       `json:"date"`
	IncreasedStats []IncreasedStat `json:"increased_stats"`
}

type UserData struct {
	Name       string     `json:"name"`
	Pfp		   string 	  `json:"pfp"`
	Activities []Activity `json:"activities"`
}
