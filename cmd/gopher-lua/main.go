package main

import (
	"encoding/json"
	"fmt"
	"github.com/yuin/gopher-lua"
	"gitlab.iglooinsure.com/axinan/backend/gadget/gopher-luar"
)

type Policy struct {
	ID         uint64 `json:"ID"`
	DisplayID  string `json:"DisplayID"`
	CoiDoc     string `json:"CoiDoc"`
	ProductKey string `json:"ProductKey"`
}

type Result struct {
	shouldSendMail bool
	data string
}

func (p *Result) SetData(d string)  {
	p.data = d
}
func (p *Result) SetSendMail(sendMail bool)  {
	p.shouldSendMail = sendMail
}

func (p *Result) ShouldSendMail() bool {
	return p.shouldSendMail
}

const script = `
print("Hello from Lua, " .. policy.DisplayID .. "!")

if policy.ProductKey == 'LAZADA1'
then
  result:SetSendMail(true)
else
  result:SetSendMail(false)
end

d = %s
print(d)
local t={}
for key, value in pairs(d) do
	print(key, value)
    table.insert(t, string.format("\"%%s\":\"%%s\"", key, value))
end
print(t)
print(t["policy"])
data = "{" .. table.concat(t, ",") .. "}"
result:SetData(data)
print('xxx')
print((true and true) ~= true)
`

func main() {
	L := lua.NewState()
	defer L.Close()

	policy := &Policy{
		DisplayID: "TEST123456",
		ProductKey: "LAZADA1",
	}
	result := &Result{}
	L.SetGlobal("policy", luar.New(L, policy))
	L.SetGlobal("result", luar.New(L, result))
	if err := L.DoString(fmt.Sprintf(script, "{policy=policy.DisplayID}")); err != nil {
		panic(err)
	}

	fmt.Println("Lua set your result to:", result.ShouldSendMail())
	fmt.Println("data", result.data)
	marshal, err := json.Marshal(policy)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}