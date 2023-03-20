package main

import (
	"fmt"
	"github.com/Shopify/go-lua"
)

func main() {
	l := lua.NewState()
	lua.OpenLibraries(l)
	err := lua.DoString(l, `function hello(x)
  if x == 1 
  then
  	return false
  else
  	return true
  end
end`)
	if err != nil {
		return 
	}
	l.Global("hello")
	err = l.ProtectedCall(0, 1, 0)
	if err != nil {
		return
	}
	x := l.ToBoolean(-1)
	fmt.Println(x)
}