package api

// IAction 动作抽象类
type IAction interface {
	GetAction() string
	GetParams() map[string]interface{}
}

// ActionStruct 动作结构体
type ActionStruct struct {
	apiUrl string
	action string
	params map[string]interface{}
}

// Build 生成要执行的动作类
func Build(apiUrl string) *ActionStruct {
	action := new(ActionStruct)
	action.apiUrl = apiUrl
	return action
}

// GetAction 获取动作名
func (a *ActionStruct) GetAction() string {
	return a.action
}

// GetParams 获取动作执行所需的参数
func (a *ActionStruct) GetParams() map[string]interface{} {
	return a.params
}
