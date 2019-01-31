package app

import (
	size "app/vendors/size/models"
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"io/ioutil"
	"math"
	"math/big"
	rand1 "math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//region 模版自定义函数，可以在模板进行调用
func TemplateFunc() template.FuncMap {
	return template.FuncMap{
		"SizeFormat": size.SizeFormat,
	}
}

//endregion

//region 自定义time.Time json输出格式
type Time time.Time

func (c Time) String() string {
	return time.Time(c).Format("2006-01-02 15:04:05")
}

//endregion

type Date time.Time

func (c Date) String() string {
	return time.Time(c).Format("2006-01-02")
}

//region MD5加密
func Strmd5(str string) string {
	w := md5.New()
	w.Write([]byte(str)) // 需要加密的字符串为
	return hex.EncodeToString(w.Sum(nil))
}

//endregion

//region 获取客户端IP
func ClientIp(c *gin.Context) string {
	remoteAddr := c.Request.RemoteAddr
	if ip := c.Request.Header.Get("XRealIP"); ip != "" {
		remoteAddr = ip
	} else if ip = c.Request.Header.Get("XForwardedFor"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

//endregion

//region 通过dns服务器8.8.8.8:80获取使用的ip
func PulicIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}

//endregion

//region Int64类型的数组去重 去0
func RemoveDuplicateInt64(list []int64) []int64 {
	var x []int64 = []int64{}
	for _, i := range list {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	return x
}

//endregion

//生成随机码
func GetRandomSalt(len int64) string {
	return GetSjCode(len)
}

func GetSjCode(len int64) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes1 := []byte(str)
	result := []byte{}
	r := rand1.New(rand1.NewSource(time.Now().UnixNano()))
	var i int64
	for i = 0; i < len; i++ {
		//result = append(result, bytes[r.Intn(len(bytes))])  bytes.Count([]byte(str),nil)-1)
		result = append(result, bytes1[r.Intn(bytes.Count(bytes1, nil)-1)])
	}
	return string(result)
}

//字符串分割数组并去空
func StrSplitRe(str string) (res []int64) {
	arr := strings.Split(str, ",")
	for _, val := range arr {
		if val != "" {
			val64, _ := strconv.ParseInt(val, 10, 64)
			res = append(res, val64)
		}
	}
	return
}

// CurPath 获取当前运行目录
func CurPath() (path string) {
	file, _ := exec.LookPath(os.Args[0])
	pt, _ := filepath.Abs(file)

	return filepath.Dir(pt)
}

func Round2(f float64, n int) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n)+"f", f)
	inst, _ := strconv.ParseFloat(floatStr, 64)
	return inst
}

//绑定旅游卡验证
func CheckLyk() {
	resp, err := http.PostForm("http://zslyapi.yytxlyw.com/User/bindPassport",
		url.Values{"mobile": {"13528837032"},
			"passport_num":  {"123456"},
			"passport_code": {"123456"},
			"real_name":     {"123456"},
			"id_number":     {"123456"},
			"gender":        {"123456"},
		})

	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(body))
}

//生成指定范围的随机数
func RandInt64(min, max int64) int64 {
	maxBigInt := big.NewInt(max)
	i, _ := rand.Int(rand.Reader, maxBigInt)
	if i.Int64() < min {
		RandInt64(min, max)
	}
	return i.Int64()
}

//生成32位随机序列
func RandNewStr(strlen int) string {
	var (
		codes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
		codeLen = len(codes)
	)
	data := make([]byte, strlen)
	rand1.Seed(time.Now().UnixNano())
	for i := 0; i < strlen; i++ {
		idx := rand1.Intn(codeLen)
		data[i] = byte(codes[idx])
	}

	return string(data)
}

//截取字符串
func SubString(str string, begin, length int) string {
	rs := []rune(str)
	lth := len(rs)
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length

	if end > lth {
		end = lth
	}
	return string(rs[begin:end])
}

//region Remark:生成csv文件并下载 Author:tang
func DownCsv(c *gin.Context, fileName string, b *bytes.Buffer) {
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", fileName))
	tet := b.String()
	c.String(200, tet)
	c.Next()
}

//endregion

//region Remark:格式化时分秒 Author:tang
func GetTimeName(agoTime int64) string {
	old_time := time.Unix(agoTime, 0)
	time_init := time.Now()
	var num int
	new_time := time_init.Unix() - old_time.Unix()
	if new_time >= 31104000 {
		num = int(new_time / 31104000)
		return strconv.Itoa(num) + "年前"
	}
	if new_time >= 2592000 {
		num = int(new_time / 2592000)
		return strconv.Itoa(num) + "月前"
	}
	if new_time >= 86400 {
		num = int(new_time / 86400)
		return strconv.Itoa(num) + "天前"
	}
	if new_time >= 3600 {
		num = int(new_time / 3600)
		return strconv.Itoa(num) + "小时前"
	}
	if new_time >= 60 {
		num = int(new_time / 60)
		return strconv.Itoa(num) + "分钟前"
	}
	return old_time.String()
}
func GetMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
func Guid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5(base64.URLEncoding.EncodeToString(b))
}

//region Remark: 获取上个月的开始时间和结束 Author tang
func LastMonthStartAndEnd() (time.Time, time.Time) {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	start := thisMonth.AddDate(0, -1, 0)
	end := thisMonth.AddDate(0, 0, -1)
	return start, end
}

//region Remark: 获取该月时间 Author tang
func GetThisMonthTime() time.Duration {
	year, month, _ := time.Now().Date()
	start := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	end := start.AddDate(0, 1, 0)
	return end.Sub(start)
}

//endregion

//region Remark:距现在时长 Author:   tang
func timeSub(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}
func FromNowTime(deal_time time.Time) string {
	toBeCharge := time.Now().Sub(deal_time)
	hour := toBeCharge.Hours()
	minutes := toBeCharge.Minutes()
	if deal_time.Format("2006-01-02") != time.Now().Format("2006-01-02") {
		return strconv.Itoa(timeSub(time.Now(), deal_time)) + " 天前"
	} else if hour > 1 && hour < 24 {
		data := math.Ceil(hour)
		return strconv.FormatFloat(data, 'f', -1, 64) + " 小时前"
	} else {
		if minutes > 0 {
			data := math.Ceil(minutes)
			return strconv.FormatFloat(data, 'f', -1, 64) + " 分钟前"
		} else {
			data := float64(1)
			return strconv.FormatFloat(data, 'f', -1, 64) + " 分钟前"
		}
	}

}

//endregion
