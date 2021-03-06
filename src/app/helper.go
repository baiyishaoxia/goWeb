package app

import (
	size "app/vendors/size/models"
	"archive/zip"
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
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
	"regexp"
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

//region   解压缩   Author:tang
func Unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}
		//------------注入

		dir := filepath.Dir(path)
		if len(dir) > 0 {
			if _, err = os.Stat(dir); os.IsNotExist(err) {
				err = os.MkdirAll(dir, 0755)
				if err != nil {
					return err
				}
			}
		}

		//---------------------end

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}
	return nil
}

//endregion
//region   下载模板   Author:tang
func TemplateeDown(c *gin.Context) {
	header := []string{"姓名(*)", "手机号码(*)", "参与活动的名称(*)"}
	b := &bytes.Buffer{}
	b.WriteString("\xEF\xBB\xBF")
	wr := csv.NewWriter(b)
	wr.Write(header) //按行shu
	wr.Flush()
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", "template.csv"))
	tet := b.String()
	c.String(200, tet)
	c.Next()
}

//endregion

//region Remark:除去文档中的样式
func RemoveHtmlStyle(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	// 去除空格
	src = strings.Replace(src, " ", "", -1)
	// 去除换行符
	src = strings.Replace(src, "\n", "", -1)
	return src
}

//endregion

func Uuid() string {
	u, _ := uuid.NewV4()
	return u.String()
}
func IsAjax(c *gin.Context) bool {
	if c.GetHeader("X-Requested-With") == "XMLHttpRequest" {
		return true
	} else {
		return false
	}
}

//说明：draw: 表示请求次数，recordsTotal: 总记录数，recordsFiltered: 过滤后的总记录数，data: 具体的数据对象数组
func TableJson(code int, count int64, msg string, data string) string {
	return "{\"draw\":" + strconv.Itoa(code) + ",\"recordsTotal\":" + strconv.FormatInt(count, 10) + ",\"data\":" + data + ",\"recordsFiltered\":\"" + msg + "\"}"
}

//字符串分割数组并去空
func StrSplitArray(str string) (res []string) {
	arr := strings.Split(str, ",")
	for _, val := range arr {
		if val != "" {
			res = append(res, val)
		}
	}
	return
}

//region Remark:除去 script
func RemoveHtmlScript(src string) string {
	//去除SCRIPT
	re, _ := regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	// 去除空格
	src = strings.Replace(src, " ", "", -1)
	// 去除换行符
	src = strings.Replace(src, "\n", "", -1)
	return src
}

//endregion

//region Remark: 字符串数组去重  tang
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
//endregion

func DateFormat(format string, timestamp ...int64) string {
	var ts = time.Now().Unix()
	if len(timestamp) > 0 {
		ts = timestamp[0]
	}
	var t = time.Unix(ts, 0)
	Y := strconv.Itoa(t.Year())
	m := fmt.Sprintf("%02d", t.Month())
	d := fmt.Sprintf("%02d", t.Day())
	H := fmt.Sprintf("%02d", t.Hour())
	i := fmt.Sprintf("%02d", t.Minute())
	s := fmt.Sprintf("%02d", t.Second())
	format = strings.Replace(format, "Y", Y, -1)
	format = strings.Replace(format, "m", m, -1)
	format = strings.Replace(format, "d", d, -1)
	format = strings.Replace(format, "H", H, -1)
	format = strings.Replace(format, "i", i, -1)
	format = strings.Replace(format, "s", s, -1)
	return format
}

/*
 *发送post请求   Author:tang
 *@param apiUrl api地址
 *@param postParam post参数
 *@param result map格式解析json数据， err error对象
 */
func UrlPost(apiUrl string, param Mp) (Mp, error) {
	if apiUrl == "" {
		return nil, errors.New("Please check the request address!")
	}
	v := url.Values{}
	for key, value := range param {
		v.Set(key, ToString(value))
	}
	u, _ := url.ParseRequestURI(apiUrl)
	urlStr := u.String()

	tr := &http.Transport{    //解决x509: certificate signed by unknown authority
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   15 * time.Second,
		Transport: tr,    //解决x509: certificate signed by unknown authority
	}
	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(v.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")      //添加请求头Type
	r.Header.Add("Content-Length", strconv.Itoa(len(v.Encode())))          //添加请求头Length
	r.Header.Add("Authorization", ToString(param["authorization"]))        //添加请求头token

	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	obj := Mp{}
	err = json.Unmarshal([]byte(body), &obj)
	if obj.Has("code") {
		if _, _code := obj.Get("code").(float64); !_code {
			return obj, errors.New(ToString(obj.Get("message")))
		}
	}
	if obj.Has("status") && (ToInt64(obj.Get("status")) != 0 && ToInt64(obj.Get("status")) != 200) {
		if obj.Has("message") {
			return obj, errors.New(ToString(obj.Get("message")))
		}
		return obj, errors.New(ToString(obj.Get("msg")))
	}
	if obj.Has("data") {
		return obj, nil
	}
	return obj, nil
}