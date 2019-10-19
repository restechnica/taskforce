package variables

type JSONCache struct {
	JSONStringMap   map[string]string
	JSONVariableMap map[string]string
}

func NewJSONCache() (cache JSONCache) {
	return JSONCache{
		JSONStringMap:   make(map[string]string),
		JSONVariableMap: make(map[string]string),
	}
}

func (cache JSONCache) AddJSON(name string, json string) {
	cache.JSONStringMap[name] = json
}

func (cache JSONCache) AddVariable(jsonName string, name string, value string) {
	var key = jsonName + ":" + name
	cache.JSONVariableMap[key] = value
}

func (cache JSONCache) GetJSON(key string) (json string, hasJSON bool) {
	json, hasJSON = cache.JSONStringMap[key]
	return
}

func (cache JSONCache) GetVariable(jsonName string, name string) (value string, hasVariable bool) {
	var key = jsonName + ":" + name
	value, hasVariable = cache.JSONVariableMap[key]
	return
}
