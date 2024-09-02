// This code is part of player package
package exerciseModel

// import other packages

type Exercise struct {
	ID         string `gorm:"index"`
	ExerciseId string `gorm:"primarykey"`
	Name       string `json:"Name"`
	Reps       string `json:"Reps"`
	Sets       string `json:"Sets"`
	BucketUrl  string `json:"BucketUrl"`
}
