package model

type Profile struct {
	ID, Name, Height, Weight, Job, JobAddress string
	Edu, Child, Jiguan, Age, Marriage         string
	Sex                                       string
	Income                                    string
}

//
//func (p Profile) String() string {
//	return fmt.Sprintf("Age: %v, Height: %v", p.Age, p.Height)
//}
