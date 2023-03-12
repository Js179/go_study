package ref

import "fmt"

type PGender uint

const (
	Male PGender = iota
	Female
)

type People struct {
	age    uint
	name   string
	gender PGender `json:"gender,omitempty,string"`
	salary float32
}

func (p *People) Salary() {
	p.salary = 20344.2
}

func (p People) Get() {
	fmt.Println("in get method...")
}

func (p People) SetParam(name string) {
	p.name = name
}
