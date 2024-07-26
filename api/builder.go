package api

type IAction interface {
	GetAction() string
	GetParams() map[string]interface{}
}

type ActionStruct struct {
	apiUrl string
	action string
	params map[string]interface{}
}

func Build(apiUrl string) *ActionStruct {
	action := new(ActionStruct)
	action.apiUrl = apiUrl
	return action
}

func (a *ActionStruct) GetAction() string {
	return a.action
}

func (a *ActionStruct) GetParams() map[string]interface{} {
	return a.params
}
