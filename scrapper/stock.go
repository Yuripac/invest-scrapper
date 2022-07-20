package scrapper

import (
	"fmt"
	"os/exec"
)

type Stock struct {
	Name  string
	Value float64
}

func (s Stock) SysNotify() {
	cmd := exec.Command("notify-send", s.Name, fmt.Sprint(s.Value))
	cmd.Run()
}
