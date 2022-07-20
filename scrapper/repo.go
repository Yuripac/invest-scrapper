package scrapper

import (
	"io/ioutil"
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type YMLRepo struct {
	mu   sync.Mutex
	File *os.File
}

type stockItem struct {
	Value float64 `yaml:"value"`
}

func (r *YMLRepo) UpdateMaxValue(stock Stock) (int, error) {
	r.mu.Lock()

	defer func() {
		r.File.Seek(0, 0)
		r.mu.Unlock()
	}()

	content, err := ioutil.ReadAll(r.File)
	if err != nil {
		return 0, err
	}

	items := make(map[string]*stockItem)
	err = yaml.Unmarshal(content, &items)
	if err != nil {
		return 0, err
	}

	updated := false
	if item := items[stock.Name]; item != nil {
		if stock.Value > item.Value {
			items[stock.Name] = &stockItem{Value: stock.Value}
			item.Value = stock.Value
			updated = true
		}
	} else {
		items[stock.Name] = &stockItem{Value: stock.Value}
		updated = true
	}

	if updated {
		newContent, err := yaml.Marshal(&items)
		if err != nil {
			return 0, err
		}
		r.File.Truncate(0)
		r.File.Seek(0, 0)

		n, err := r.File.Write(newContent)
		if err != nil {
			return 0, err
		}
		log.Printf("%s updated to %.2f", stock.Name, stock.Value)
		return n, nil
	}
	return 0, nil
}
