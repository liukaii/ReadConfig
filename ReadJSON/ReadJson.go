package ReadJSON

import "os"

type Config struct {
	MyMap map[string]string
}

func (*Config) InitConfig(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

}
