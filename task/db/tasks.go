package db

import (
	"encoding/binary"
	// go get -u github.com/boltdb/bolt/...
	"github.com/boltdb/bolt"
	"time"
)

var tasksBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(tasksBucket)
		return err
	})
}

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := iTob(int(id64))
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func AllTasks() ([]Task, error) {
	var tasks []Task

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask (key int) error{
	return db.Update(func(tx *bolt.Tx) error {
		b:=tx.Bucket(tasksBucket)
		return b.Delete(iTob(key))
	})
}

func iTob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
