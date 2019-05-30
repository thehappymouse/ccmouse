package engine

import (
	"github.com/thehappymouse/go-utils"
	"encoding/json"
	"io/ioutil"
	"sync"
	"github.com/rs/zerolog/log"
)

type JsonStore struct {
	source   map[string]interface{}
	JsonPath string
	once     sync.Once
	sync.RWMutex
}

func (t *JsonStore) Set(key string, g interface{}) bool {
	//t.once.Do(func() {
	//	t.LoadDisk()
	//})
	t.Lock()
	defer t.Unlock()
	t.source[key] = g
	return true
}

func (t *JsonStore) Get(id string) interface{} {
	t.RLock()
	defer t.RUnlock()

	v, ok := t.source[id]
	if !ok {
		return nil
	}
	return v
}

// 写盘
func (t *JsonStore) WriteDisk() int {
	t.Lock()
	defer t.Unlock()
	bytes, err := json.MarshalIndent(t.source, "  ", "  ")
	utils.CheckError(err)
	utils.CheckError(ioutil.WriteFile(t.JsonPath, bytes, 0777))
	return len(t.source)
}

// 读盘
func (t *JsonStore) LoadDisk() bool {
	t.Lock()
	defer t.Unlock()
	bytes, err := ioutil.ReadFile(t.JsonPath)
	t.source = make(map[string]interface{})
	if err != nil {
		log.Error().Msgf("加载[%s]数据失败:[%s]", t.JsonPath, err)
		return false
	}
	utils.CheckError(json.Unmarshal(bytes, &t.source))
	return true
}

func CreateJsonStore(jsonPath string) *JsonStore {
	t := JsonStore{JsonPath: jsonPath}
	t.LoadDisk()
	return &t
}
