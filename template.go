package common

import (
	"bufio"
	"bytes"
	"html/template"
	"strings"
)

type Template struct {
	Name string      // 模版名字
	Data interface{} // 原始结构体
	Text string      // 模版文件
}

func NewTemplate(name string, data interface{}, text string) *Template {
	return &Template{
		Name: name,
		Data: data,
		Text: text,
	}
}

func unescaped(str string) template.HTML { return template.HTML(str) }

func (t *Template) ParseTpl() (tplData string, err error) {
	var (
		optSlice []string
		buf      bytes.Buffer
	)

	tpl := template.New(t.Name)
	tpl = tpl.Funcs(template.FuncMap{"unescaped": unescaped})
	parse, err := tpl.Parse(t.Text)
	if err != nil {
		return
	}

	if err = parse.Execute(&buf, t.Data); err != nil {
		return
	}
	reader := bytes.NewReader(buf.Bytes())
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			optSlice = append(optSlice, line)
		}
	}
	tplData = strings.Join(optSlice, "\n")
	return
}
