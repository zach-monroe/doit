package todo

import (
	"encoding/json"
	"os"
)

type Item struct {
	Text string
}

func SaveItems(filename string, items []Item) error {
	a, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	if a != nil {
		var orig_items []Item
		if err := json.Unmarshal(a, &orig_items); err != nil {
			return err
		}
		items = append(orig_items, items...)
	}

	b, err := json.Marshal(items)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}
	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}
	return items, nil

}
