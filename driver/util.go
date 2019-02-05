package driver

import (
	"bufio"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func Require(s string) {
	if viper.Get(s) == nil {
		if v, ok := os.LookupEnv(s); ok == false || v == "" {
			ans := Prompt(" please set required value: " + s)
			viper.SetDefault(s, ans)
			_ = os.Setenv(s, strings.ToUpper(ans))
		}
	}
}

func Prompt(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return text
}
