package models

import "sync"

type Database struct{
	URL		string		`json:"url"`
	IsDead	bool
	mutex	sync.RWMutex
}

func (database *Database) SetDead(b bool) {
	database.mutex.Lock()
	database.IsDead = b
	database.mutex.Unlock()
}

func (database *Database) GetIsDead() bool {
	database.mutex.RLock()
	isAlive := database.IsDead
	database.mutex.RUnlock()
	return isAlive
}