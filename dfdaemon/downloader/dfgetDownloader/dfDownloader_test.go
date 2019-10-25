package dfgetDownloader

import (
	"testing"

	"github.com/dragonflyoss/Dragonfly/dfget/config"
)

func TestDFGetter_Download(t *testing.T) {
	c := config.NewConfig()
	getter := NewGetter(c)
	getter.Download("www.google.com",  nil, "test")
}
