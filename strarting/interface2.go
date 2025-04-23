package main

import "fmt"

func main(){
	fmt.Println("Level two....")
}

type Playable interface{
	Play() string
}
type guitar struct{
	plays string
}
type piano struct{
	plays string
}
type drum struct{
	plays string
}

func (g guitar) Play(){
	fmt.Println("It plays",g.plays)
}
func (p piano) Play(){
	fmt.Println("It plays", p.plays)
}
func (d drum) Play(){
	fmt.Println("It plays",d.plays)
}
func Playconcert([]Playable){
	for _,v := range
}