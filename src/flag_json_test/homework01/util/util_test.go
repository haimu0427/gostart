package util

import (
	"testing"
)

// 编写测试用例
func TestMonsterStoreAndRestore(t *testing.T) {
	var monster = Monster{
		Name:  "Goblin",
		Age:   200,
		Skill: "Stealth",
	}
	monster.Store()
	t.Log("Monster stored successfully")
	monster.Restore()
	t.Log("Monster restored successfully")
}
