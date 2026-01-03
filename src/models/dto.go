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
		return errors.New("Must have at least one 'increased_stats'")
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
	Title          *string          `json:"title,omitempty"`
	Description    *string          `json:"description,omitempty"`
	Date           *string          `json:"date,omitempty"`
	IncreasedStats *[]IncreasedStat `json:"increased_stats,omitempty"`
}

func (value ActivityModifyInput) Validate() error {
	if value.IncreasedStats != nil && len(*value.IncreasedStats) == 0 {
		return errors.New("Must have at least one 'increased_stats' if provided")
	}

	if value.IncreasedStats != nil {
		for _, v := range *value.IncreasedStats {
			err := v.Validate()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
