package stringutils

import {
  "fmt"
  "strings"
}

type StringUtils struct {

}

func Upper(val string) (string) {
  return strings.ToUpper(val)
}

func Lower(val string) (string) {
  return strings.ToLower(val)
}
