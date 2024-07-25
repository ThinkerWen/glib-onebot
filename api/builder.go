package api

type IAction interface {
	GetAction() string
	GetParams() map[string]interface{}
}

type ActionStruct struct {
	action string
	params map[string]interface{}
}

func Build() *ActionStruct {
	return new(ActionStruct)
}

func (a *ActionStruct) GetAction() string {
	return a.action
}

func (a *ActionStruct) GetParams() map[string]interface{} {
	return a.params
}
