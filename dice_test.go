package dice

import (
	"fmt"
	"testing"
)

func TestSides(t *testing.T) {
	d := die{faces: []int{1, 2, 3, 4, 5, 6}}
	for i := 0; i < len(d.faces); i++ {
		if d.side(i) != i+1 {
			t.Errorf("Expected: 0, got %d", d.side(0))
		}
	}
}

func TestNewDie(t *testing.T) {
	d := NewDie(20)
	for i := 0; i < len(d.faces); i++ {
		if d.side(i) != i+1 {
			t.Errorf("Expected: 0, got %d", d.side(0))
		}
	}
}

func TestRollOne(t *testing.T) {
	d := NewDie(6)
	for i := 0; i < 100; i++ {
		r := d.Roll()
		if !(r >= 1 && r <= 6) {
			t.Error("roll out of range")
		}
	}
}

func TestRollTwo(t *testing.T) {
	d1, d2 := NewDie(6), NewDie(6)
	for i := 0; i < 100; i++ {
		r1, r2 := d1.Roll(), d2.Roll()
		r := r1 + r2
		if !(r >= 2 && r <= 12) {
			t.Error("roll out of range", r)
		}
	}
}

func TestHardWay(t *testing.T) {
	sides := 6
	for i := 0; i < sides; i++ {
		b, s := IsHardWay(i, i)
		if b != true && s != fmt.Sprintf("%d the hard way!", i+i) {
			t.Error("not the hard way", i)
		}
	}
	for i := 0; i < sides; i++ {
		j := i + 1
		b, _ := IsHardWay(i, j)
		if b == true {
			t.Error("not the hard way", i, j)
		}
	}
}

func TestProb(t *testing.T) {
	var graph = make(map[int]int)
	sides := 6

	d1, d2 := NewDie(sides), NewDie(sides)
	for i := 0; i < 10000; i++ {
		r1, r2 := d1.Roll(), d2.Roll()
		graph[r1+r2] += 1
	}

	var sum int
	var pct float64

	fmt.Println("Results of 10,000 throws")
	for i := 2; i <= 12; i++ {
		sum += graph[i]
		pct += (float64(graph[i]) / float64(10000)) * 100
		fmt.Printf("%2d : %4d : %5.2f%%\n", i, graph[i], (float64(graph[i])/float64(10000))*100)
	}
	fmt.Printf("    %4d :%6.2f%%\n", sum, pct)
}
