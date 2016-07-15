package logfile

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/c4pt0r/ini"
)

const FILE_NO int = 25

type FileModel struct {
	BasePath *string
	Mutexes  map[int]*sync.Mutex
}

func (this *FileModel) Init() {

	conf := ini.NewConf("base.ini")
	this.BasePath = conf.String("sys", "path", "data/")
	conf.Parse()

	this.Mutexes = make(map[int]*sync.Mutex)
	for i := 0; i < FILE_NO; i++ {
		this.Mutexes[i] = &sync.Mutex{}
	}
}

func (this *FileModel) Record(content []byte) {
	dirPath := *this.BasePath + fmt.Sprintf("%d%d%d%d", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour())
	err := os.Chdir(dirPath)
	if err != nil {
		os.Mkdir(dirPath, 0755)
	}

	fileNo := fmt.Sprintf("%d-%d-%d-%d", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour())
	rSource := rand.New(rand.NewSource(time.Now().Unix()))
	key := rSource.Intn(FILE_NO)

	fileName := dirPath + "/" + fileNo + strconv.Itoa(key) + ".data"
	this.Mutexes[key].Lock()
	defer this.Mutexes[key].Unlock()
	fileObj, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	defer fileObj.Close()
	if err == nil {
		fileObj.WriteString(string(bytes.TrimRight(content, "\x00")) + "\n")
	}
}
