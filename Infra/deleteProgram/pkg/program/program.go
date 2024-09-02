// This code is part of player package
package programModel

// import other packages
import (
	exerciseModel "deleteProgram/main/pkg/exercise"

	"gorm.io/gorm"
)

type Program struct {
	gorm.Model
	ProgramId   string                   `gorm:"index,primarykey"`
	ClientId    string                   `json:"ClientId"`
	ProgramName string                   `json:"ProgramName"`
	Description string                   `json:"Description"`
	Frequency   string                   `json:"Frequency"`
	Duration    string                   `json:"Duration"`
	IsCompleted string                   `json:"IsCompleted"`
	Exercises   []exerciseModel.Exercise `gorm:"foreignKey:ID"`
	//ExericseId string `json:"ExerciseId"`
}
