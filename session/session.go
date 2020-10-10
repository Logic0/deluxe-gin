package session

import (
    "crypto/md5"
    "encoding/hex"
    "math/rand"
    "strconv"
    "strings"
    "time"
)

type Session struct{
    SessionID string          `json:"session_id"`
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
    letterIdxBits = 6                    // 6 bits to represent a letter index
    letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
    letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMask( n int ) string {
    sb := strings.Builder{}
    sb.Grow(n)
    // A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
    for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
        if remain == 0 {
            cache, remain = src.Int63(), letterIdxMax
        }
        if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
            sb.WriteByte(letterBytes[idx])
            i--
        }
        cache >>= letterIdxBits
        remain--
    }

    return sb.String()
}

// 生成32位MD5
func MD5(text string) string{
    ctx := md5.New()
    ctx.Write([]byte(text))
    return hex.EncodeToString(ctx.Sum(nil))
}

func GenerateSessionID( appid uint, id string ) string{
    //return RandStringBytesMask(36 )
    return "txop" + MD5( strconv.FormatUint( uint64( appid ), 10 ) + id )
}
