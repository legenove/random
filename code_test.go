package random

import (
	"fmt"
	"sync"
	"github.com/legenove/utils"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func str2Intstr(s string) string {
	var n = ""
	for _, one := range s {
		n += fmt.Sprintf("%d", one)
	}
	return n
}

func GetRandomString(str string) string {
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(str); i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}


//append s with b until len(s) >= l
func strPad(s, b string, l int) string {
	for len(s) < l {
		s += b
	}
	return s
}


//year month day: 20190102
func getYearMonthDay(format string) string {
	t := time.Now().Format(format)
	return t
}

func CreateCode(location string) (string, error) {
	var code string
	id := UuidV5()
	idstr := str2Intstr(id)
	fmt.Println(idstr, len(idstr))
	lenth := len(idstr)
	a := GetRandomString(idstr)
	fmt.Println(a)
	idstr = a[lenth-10:]
	fmt.Println(idstr)

	env := strings.ToUpper(location)
	index := strings.IndexByte(env, '_')
	local := env[index+1 : index+3]

	code = fmt.Sprintf("%s%s%s", local,utils.GetCurrentFormat("20060102"), idstr)
	code = strPad(code, "0", 20)
	return code[:20], nil
}

func TestCreateCode(t *testing.T) {
	for i:= 0; i < 10 ; i++ {
		fmt.Println(CreateCode("em_Ab"))
	}

}

func TestChannelRandom(t *testing.T) {
	for i:= 0; i < 10 ; i++ {
		fmt.Println(<-ch)
	}
}


func TestUniqid(t *testing.T) {
	for i:= 19; i <33; i ++ {
		s := UniqStringId(i, )
		fmt.Println(s, len(s))
	}
}


func TestXid(t *testing.T) {
	wg := sync.WaitGroup{}
	for j:= 0; j <10 ; j ++ {
		wg.Add(1)
		go func() {
			for i:= 0; i < 10 ; i++ {
				a := Xid()
				fmt.Println(a, len(a))
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
