package dice

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

type die struct {
	faces []int
}

func (d die) side(f int) int {
	return d.faces[f]
}

func NewDie(s int) (d die) {
	d = die{}
	d.faces = make([]int, s)
	for s > 0 {
		d.faces[s-1] = s
		s--
	}
	return
}

func (d die) Roll() (r int) {
	roll, err := getRand(len(d.faces))
	if err != nil {
		log.Fatalln("rng broke", err)
	}
	r = d.side(roll)
	return
}

func getRand(l int) (r int, err error) {
	faces := big.NewInt(int64(l))
	up, err := rand.Int(rand.Reader, faces)
	if err != nil {
		return
	}
	r = int(up.Int64())
	return
}

func IsHardWay(d1, d2 int) (b bool, s string) {
	if d1 == d2 {
		b = true
		switch d1 + d2 {
		case 2:
			s = fmt.Sprintf("%d! Snake Eyes!", d1+d2)
		case 12:
			s = fmt.Sprintf("%d! Boxcars!", d1+d2)
		default:
			s = fmt.Sprintf("%d the hard way!", d1+d2)
		}
	}
	return
}
