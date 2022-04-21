package gamedata

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/recordfile"
)

func readRf(st interface{}) *recordfile.RecordFile {
	rf, err := recordfile.New(st)
	if err != nil {
		log.Fatal("%v", err)
	}
	err = rf.Read("/Users/liuyang/Desktop/Git/Go_Dev/Leaf-Chat/Servers/gamedata/robots.txt")
	if err != nil {
		log.Fatal("%v: %v", err)
	}

	return rf
}
