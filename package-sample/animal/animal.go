package animal

import "fmt"

type Bird struct {
	Name     string
	isFlying bool
}

func (b *Bird) Fly(isFlying bool){
	b.isFlying = true
}

func (b *Bird) Show(){
	if b.isFlying{
		fmt.Printf("%s is flying at the moment\n", b.Name)
		return
	}
	fmt.Printf("%s is resting now", b.Name)
}
