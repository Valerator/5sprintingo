package trainings

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("некорректный формат данных")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("некорректный формат данных")
	}
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("некорректный формат данных")
	}
	t.Steps = steps
	t.TrainingType = parts[1]
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию

}
