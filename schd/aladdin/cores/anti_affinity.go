package cores

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

/************************************************************************************************************
 *
 * @copyright Institute of Software, CAS
 * @author    wuheng@iscas.ac.cn
 * @since     2018-10-16
 *
 **************************************************************************************************************/

type AntiAffinity struct {
	values    map[string]int
}

func NewAntiAffinity() *AntiAffinity {
	return &AntiAffinity{
		values:    make(map[string]int),
	}
}

func (aa *AntiAffinity) ReadLine(fileName string) (map[string]int, error) {
	f, err := os.Open(fileName)

	if err != nil {
		return aa.values, err
	}

	buf := bufio.NewReader(f)

	for {

		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)

		if err != nil && err == io.EOF {
			return  aa.values, nil
		}

		str := strings.Split(line, ",")
		if len(str[0]) == 0 || len(str[1]) == 0 || len(str[2]) == 0 {
			log.Print(str)
		} else {
			value, err := strconv.Atoi(str[2])
			if err == nil {
				aa.values[AntiAffinityKey(str[0], str[1])] = value
			}
		}
	}

	return aa.values, err
}

func (aa *AntiAffinity) GetAntiAffinity() map[string]int  {
	return aa.values
}
