package models

import "errors"

type ActivityInput struct {
	Title          string          `json:"title"`
	Description    string          `json:"description"`
	IncreasedStats []IncreasedStat `json:"increased_stats"`
}

func (value ActivityInput) Validate() error {
	if value.Title == "" {
		return errors.New("Field 'title' is required")
	}
	if len(value.IncreasedStats) == 0 {
		return errors.New("At least one stat must be provided")
	}

	for _, v := range value.IncreasedStats {
		err := v.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

type ActivityModifyInput struct {
	Title          string          `json:"title,omitempty"`
	Description    string          `json:"description,omitempty"`
	Date           string          `json:"date,omitempty"`
	IncreasedStats []IncreasedStat `json:"increased_stats,omitempty"`
}

func (value ActivityModifyInput) Validate() error {
	if value.Title == "" {
		return errors.New("Field 'title' is required")
	}
	if value.Date == "" {
		return errors.New("Field 'date' is required")
	}
	if len(value.IncreasedStats) == 0 {
		return errors.New("At least one stat must be provided")
	}

	if value.IncreasedStats != nil {
		for _, v := range value.IncreasedStats {
			err := v.Validate()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type Stats struct {
	Knowledge   int
	Guts        int
	Proficiency int
	Kindness    int
	Charm       int
}

type HomePageData struct {
	UserData
	Stats Stats `json:"stats"`
}

type Username struct {
	Name string
}

type Modal struct {
	Title   string
	Content string
	Data 	any
}

type Toast struct {
	Type    string
	Message string
}
