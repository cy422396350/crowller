package itemServer

import (
	"github.com/cy422396350/crowller/zhenai/model"
	"testing"
)

func TestSave(t *testing.T) {
	profile := model.Profile{
		Name:      "冰之泪",
		Age:       47,
		Income:    "8001-12000元",
		Education: "大专",
	}
	save(profile)
}
