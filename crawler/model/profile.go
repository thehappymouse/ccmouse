package model

import "fmt"

type Profile struct {
	ID, Name, Heieght, Weight, Job, JobAddress string
	Age                                        int
}

func (p Profile) String() string {
	return fmt.Sprintf("Age: %d", p.Age)
}
