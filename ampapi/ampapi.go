package ampapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type AMPAPI struct {
	baseURI    string
	sessionId  string
	dataSource string
	APISpec    map[string]map[string]map[string]interface{}
	API        map[string]map[string]func([]interface{}) string
}

func NewAMPAPI(baseURI string) *AMPAPI {
	baseUri := baseURI
	var dataSource string = ""
	if string(baseUri[len(baseUri)-1]) == "/" {
		dataSource = baseUri + "API"
	} else {
		dataSource = baseUri + "/API"
	}

	return &AMPAPI{
		baseURI:    baseURI,
		sessionId:  "",
		dataSource: dataSource,
		APISpec:    map[string]map[string]map[string]interface{}{"Core": {"GetAPISpec": {}}},
		API:        map[string]map[string]func([]interface{}) string{},
	}
}

// APICall is a function that takes an endpoint and returns the response as a string
func (a *AMPAPI) aPICall(module string, methodName string, args []interface{}) string {
	data := map[string]interface{}{}
	method := a.APISpec[module][methodName]

	val, ok := method["Parameters"]
	if ok {
		methodParams := val.([]interface{})

		for i := 0; i <= len(methodParams)-1; i++ {
			key := methodParams[i].(map[string]interface{})["Name"].(string)
			data[key] = args[i]
		}
	}

	data["SESSIONID"] = a.sessionId

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(data)

	client := &http.Client{}
	req, err := http.NewRequest("POST", a.dataSource+"/"+module+"/"+methodName, body)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	return string(result)
}

func (a *AMPAPI) Login(username string, password string, token string, rememberMe bool) map[string]interface{} {
	if a.APISpec["Core"]["Login"] == nil {
		res := a.aPICall("Core", "GetAPISpec", nil)
		var result map[string]map[string]map[string]map[string]interface{}
		err := json.Unmarshal([]byte(res), &result)

		if err != nil {
			fmt.Print(err.Error())
			return nil
		}

		a.APISpec = result["result"]
	}

	res := a.aPICall("Core", "Login", []interface{}{username, password, token, rememberMe})

	var loginResult map[string]interface{}
	err := json.Unmarshal([]byte(res), &loginResult)
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}

	if loginResult["success"].(bool) {
		a.sessionId = loginResult["sessionID"].(string)
	}

	return loginResult
}

func (a *AMPAPI) Call(module string, method string, args []interface{}) string {
	if a.APISpec[module] == nil || a.APISpec[module][method] == nil {
		res := a.aPICall("Core", "GetAPISpec", nil)
		var result map[string]map[string]map[string]map[string]interface{}
		err := json.Unmarshal([]byte(res), &result)

		if err != nil {
			fmt.Print(err.Error())
			return "{\"success\":false}"
		}

		a.APISpec = result["result"]
	}

	res := a.aPICall(module, method, args)

	return res
}

// func (a *AMPAPI) Init(stage2 bool) bool {
// 	for module := range a.APISpec {
// 		for method := range a.APISpec[module] {
// 			if a.API[module] == nil {
// 				a.API[module] = map[string]func([]interface{}) string{}
// 			}
// 			a.API[module][method] = func(args []interface{}) string {
// 				m := Method{
// 					module: module,
// 					method: method,
// 				}

// 				return a.APICall(m, args)
// 			}

// 			fmt.Println(module)
// 			fmt.Println(method)
// 			fmt.Println(a.API[module][method])
// 		}
// 	}

// 	if stage2 {
// 		return true
// 	} else {
// 		res := a.API["Core"]["GetAPISpec"](nil)

// 		var result map[string]map[string]map[string]map[string]interface{}
// 		err := json.Unmarshal([]byte(res), &result)
// 		if err != nil {
// 			fmt.Print(err.Error())
// 			return false
// 		}

// 		if (result != nil) && (result["result"] != nil) {
// 			a.APISpec = result["result"]
// 			return a.Init(true)
// 		} else {
// 			return false
// 		}
// 	}
// }
