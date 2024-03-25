package conf_test

import (
	"os"
	"testing"

	"github.com/liuli404/VBlog/conf"
)

func TestLoadFromFile(t *testing.T) {
	err := conf.LoadFromFile("./application.toml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C())
}

func TestLoadFromEnv(t *testing.T) {
	// os.Setenv("DATASOURCE_HOST", "110.41.160.251")
	os.Setenv("DATASOURCE_PORT", "3310")
	err := conf.LoadFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C())
}
