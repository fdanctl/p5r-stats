package services

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/fdanctl/p5r-stats/config"
	"github.com/fdanctl/p5r-stats/src/models"
	"github.com/fdanctl/p5r-stats/src/utils"
)

func writeData(userData *models.UserData) error {
	data, err := json.MarshalIndent(&userData, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(config.DataFile, data, 0o644); err != nil {
		return err
	}

	return nil
}

func ComputeStats(activities []models.Activity) map[models.Stat]int {
	result := make(map[models.Stat]int, 5)
	result[models.Knowledge] = 0
	result[models.Guts] = 0
	result[models.Proficiency] = 0
	result[models.Kindness] = 0
	result[models.Charm] = 0

	for _, v := range activities {
		for _, s := range v.IncreasedStats {
			result[s.Stat] += int(s.Points)
		}
	}

	return result
}

func ReadUserData() (*models.UserData, error) {
	file, err := os.ReadFile(config.DataFile)
	if err != nil {
		return nil, models.ErrNotFound
	}

	var data models.UserData
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func CreateUserData(name string) (*models.UserData, error) {
	_, err := os.ReadFile(config.DataFile)

	if err == nil {
		return nil, models.ErrAlreadyExists
	}

	newData := models.UserData{
		Name:       name,
		Pfp:        "default_pfp.png",
		Activities: (make([]models.Activity, 0)),
	}

	err = writeData(&newData)
	if err != nil {
		return nil, err
	}

	return &newData, nil
}

func ModifyUser(name string, fh *multipart.FileHeader) error {
	userData, err := ReadUserData()
	var path *string
	if err != nil {
		return err
	}

	if fh != nil {
		file, err := fh.Open()
		if err != nil {
			return models.ErrCantReadFile
		}
		defer file.Close()

		outPath := filepath.Join(
			"assets",
			fmt.Sprint(name, "_pfp", filepath.Ext(fh.Filename)),
		)

		outFile, err := os.Create(outPath)
		if err != nil {
			return err
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, file)
		if err != nil {
			return err
		}
		path = &outPath
	}

	userData.Name = utils.FallbackToB(&name, &userData.Name)
	userData.Pfp = utils.FallbackToB(path, &userData.Pfp)

	writeData(userData)
	return nil
}

func DeleteUserData() error {
	err := os.Remove(config.DataFile)
	return err
}

func ReadActivity(id string) (*models.Activity, error) {
	file, err := os.ReadFile(config.DataFile)
	if err != nil {
		return nil, models.ErrNotFound
	}

	var data models.UserData
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	for _, v := range data.Activities {
		if v.Id == id {
			return &v, nil
		}
	}

	return nil, models.ErrNotFound
}

func InsertActivity(input models.ActivityInput) (*models.Activity, error) {
	userData, err := ReadUserData()
	if err != nil {
		return nil, err
	}

	randId, err := utils.RandomID(4)
	if err != nil {
		return nil, err
	}

	activity := models.Activity{
		Id:             randId,
		Title:          input.Title,
		Description:    input.Description,
		Date:           time.Now(),
		IncreasedStats: input.IncreasedStats,
	}
	userData.Activities = append(userData.Activities, activity)

	writeData(userData)

	return &activity, nil
}

func ModifyActivity(id string, input models.ActivityModifyInput) error {
	userData, err := ReadUserData()
	if err != nil {
		fmt.Println("Data does not exists")
		return err
	}

	var found bool
	for i, v := range userData.Activities {
		if v.Id == id {
			var pDate *time.Time
			if *input.Date != "" {
				layout := "2006-01-02"
				inputDate, err := time.Parse(layout, *input.Date)
				if err != nil {
					return models.ErrInvalidDate
				}
				pDate = &inputDate
			}

			userData.Activities[i] = models.Activity{
				Id:             v.Id,
				Title:          utils.FallbackToB(input.Title, &v.Title),
				Description:    utils.FallbackToB(input.Description, &v.Description),
				Date:           utils.FallbackToB(pDate, &v.Date),
				IncreasedStats: utils.FallbackToB(input.IncreasedStats, &v.IncreasedStats),
			}
			found = true
			break
		}
	}

	if !found {
		return models.ErrNotFound
	}

	err = writeData(userData)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	return nil
}

func DeleteActivity(id string) error {
	userData, err := ReadUserData()
	if err != nil {
		fmt.Println("Data does not exists")
		return err
	}

	var targetActivity models.Activity
	for i, v := range userData.Activities {
		if v.Id == id {
			userData.Activities = append(
				userData.Activities[:i],
				userData.Activities[i+1:]...,
			)
			targetActivity = v
			break
		}
	}

	if len(targetActivity.Id) == 0 {
		return models.ErrNotFound
	}

	err = writeData(userData)
	if err != nil {
		return err
	}

	return nil
}
