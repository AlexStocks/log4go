package log4go

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestLogger_LoadConfiguration(t *testing.T) {
	pwd, _ := filepath.Abs(".")
	t.Logf("pwd %s", pwd)

	xmlConfFile := pwd + "/examples/example.xml"
	xmlConf := loadConfFile(xmlConfFile)
	t.Logf("xml  conf:%+v", xmlConf)

	jsonConfFile := pwd + "/examples/example.json"
	jsonConf := loadConfFile(jsonConfFile)
	t.Logf("json conf:%+v", jsonConf)

	if !reflect.DeepEqual(xmlConf, jsonConf) {
		t.Fatalf("xmlConf %+v != jsonConf %+v", xmlConf, jsonConf)
	}

	yamlConfFile := pwd + "/examples/example.yml"
	yamlConf := loadConfFile(yamlConfFile)
	t.Logf("yaml conf:%+v", yamlConf)

	if !reflect.DeepEqual(xmlConf, yamlConf) {
		t.Fatalf("xmlConf %+v != yamlConf %+v", xmlConf, yamlConf)
	}
}
