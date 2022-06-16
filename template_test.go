package common

import (
	"fmt"
	"testing"
)

const yamlTpl = `apiVersion: apps/v1
kind: Deployment
        - name: {{ .ProjectName }}
          image: {{ .Image }}
          command:
            - sh
            - -c
            - {{ .StartCommand | unescaped }}`

func TestParseTpl(t *testing.T) {
	data := struct {
		ProjectName  string
		Image        string
		StartCommand string
	}{
		"zhengyansheng",
		"100",
		"nginx && sleep 2000",
		//template.HTML("nginx && sleep 2000"),
	}
	tpl := Template{"deployment", data, yamlTpl}
	data1, err := tpl.ParseTpl()
	fmt.Println(data1)
	fmt.Println(err)
}
