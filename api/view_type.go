package api

import (
	"encoding/json"
	"fmt"
)

type ViewType int

const (
	Full ViewType = iota + 1
	Tall
	Compact
)

func (v ViewType) String() string {
	switch v {
	case Full:
		return "full"
	case Tall:
		return "tall"
	case Compact:
		return "compact"
	default:
		return "unknown"
	}
}

func NewViewType(viewType string) (ViewType, error) {
	names := [...]string{
		"full",
		"tall",
		"compact",
	}
	for i, v := range names {
		if viewType == v {
			return ViewType(i + 1), nil
		}
	}
	return 0, fmt.Errorf("err %v is invalid argment", viewType)
}

// Unmarshal時の動作を定義します
func (v *ViewType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("data should be a string, got %s", data)
	}

	f, err := NewViewType(s)
	if err != nil {
		return err
	}
	*v = f
	return nil
}

// Marshal時の動作を定義します
func (v ViewType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}
