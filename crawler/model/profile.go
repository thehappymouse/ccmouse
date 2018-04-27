package model

import "encoding/json"

type Profile struct {
	ID, Name, Height, Weight, Job, JobAddress string
	Edu, Child, Jiguan, Age, Marriage         string
	Sex                                       string
	Income                                    string
}
// todo Map 转 Struce
func Map2Profile(o interface{}) Profile  {
	str, _ := json.Marshal(o)
	p := Profile{}
	json.Unmarshal(str, &p)
	return p
}
//
//func (p Profile) String() string {
//	return fmt.Sprintf("Age: %v, Height: %v", p.Age, p.Height)
//}
