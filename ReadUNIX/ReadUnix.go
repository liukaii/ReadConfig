package ReadUNIX

import (
	"os"
	"bufio"
	"io"
	"strings"
)

type Config struct {
	Mymap	map[string]string
	strcet string
}

func (c *Config) InitConfig(path string) {
	c.Mymap = make(map[string]string)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		s := strings.TrimSpace(string(b))//去除空格
		if strings.Index(s, "#") == 0 {//判断此行是不是注释
			continue
		}

		n1 := strings.Index(s,"[")//取括号中的内容
		n2 := strings.LastIndex(s, "]")
		if n1 > -1 && n2 > -1 && n2 > n1 + 1 {
			c.strcet = strings.TrimSpace(s[n1+1:n2])
			continue
		}


		if len(c.strcet) == 0 {
			continue
		}
		index := strings.Index(s,"=")//获取等号的位置
		if index < 0 {
			continue
		}

		first := strings.TrimSpace(s[:index])//获取key值
		if len(first) == 0 {
			continue
		}
		second := strings.TrimSpace(s[index+1:])//获取value值

		pos := strings.Index(second, "\t#")
		if pos > -1 {//判断此行含不含注释
			second = second[0:pos]
		}

		pos = strings.Index(second, " #")
		if pos > -1 {//判断此行含不含注释
			second = second[0:pos]
		}

		pos = strings.Index(second, "\t//")
		if pos > -1 {//判断此行含不含注释
			second = second[0:pos]
		}

		pos = strings.Index(second, " //")
		if pos > -1 {//判断此行含不含注释
			second = second[0:pos]
		}

		if len(second) == 0 {
			continue
		}

		key := first
		c.Mymap[key] = strings.TrimSpace(second)
	}
}

func (c Config) Read(key string) string {
	v, found := c.Mymap[key]
	if !found {
		return ""
	}
	return v
}




