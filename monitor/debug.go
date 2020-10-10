package monitor

import (
    "encoding/base64"
    "fmt"
    "strconv"
    "strings"

    "github.com/pkg/errors"
)

type SpanContext struct {
    TraceID                 []int64
    ParentSegmentID         []int64
    ParentSpanID            int32
    ParentServiceInstanceID int32
    EntryServiceInstanceID  int32
    NetworkAddressID        int32
    EntryEndpointID         int32
    ParentEndpointID        int32
    Sample                  int8
    NetworkAddress          string
    EntryEndpoint           string
    ParentEndpoint          string
}
func stringConvertGlobalID(str string) ([]int64, error) {
    idStr, err := base64.StdEncoding.DecodeString(str)
    if err != nil {
        return nil, errors.Wrapf(err, "decode id error %s", str)
    }
    ss := strings.Split(string(idStr), ".")
    if len(ss) < 3 {
        return nil, errors.Errorf("decode id entities error %s", string(idStr))
    }
    ii := make([]int64, len(ss))
    for i, s := range ss {
        ii[i], err = strconv.ParseInt(s, 0, 64)
        if err != nil {
            return nil, errors.Wrapf(err, "convert id error convert id entities to int32 error %s", s)
        }
    }
    return ii, nil
}

func stringConvertInt32(str string) (int32, error) {
    i, err := strconv.ParseInt(str, 0, 32)
    return int32(i), err
}

func decodeBase64(str string) (string, int32, error) {
    ret, err := base64.StdEncoding.DecodeString(str)
    if err != nil {
        return "", 0, err
    }
    retStr := string(ret)
    if strings.HasPrefix(retStr, "#") {
        return retStr[1:], 0, nil
    }
    i, err := strconv.ParseInt(retStr, 0, 32)
    if err != nil {
        return "", 0, err
    }
    return "", int32(i), nil
}

func DecodeSW6(header string) error {
    var tc SpanContext
    if header == "" {
        return errors.New("empty header")
    }
    hh := strings.Split(header, "-")
    if len(hh) < 7 {
        return errors.New("not enough")
    }
    sample, err := strconv.ParseInt(hh[0], 10, 8)
    if err != nil {
        return errors.Errorf("str to int8 error %s", hh[0])
    }
    tc.Sample = int8(sample)
    tc.TraceID, err = stringConvertGlobalID(hh[1])
    if err != nil {
        return errors.Wrap(err, "trace id parse error")
    }
    tc.ParentSegmentID, err = stringConvertGlobalID(hh[2])
    if err != nil {
        return errors.Wrap(err, "parent segment id parse error")
    }
    tc.ParentSpanID, err = stringConvertInt32(hh[3])
    if err != nil {
        return errors.Wrap(err, "parent span id parse error")
    }
    tc.ParentServiceInstanceID, err = stringConvertInt32(hh[4])
    if err != nil {
        return errors.Wrap(err, "parent service instance id parse error")
    }
    tc.EntryServiceInstanceID, err = stringConvertInt32(hh[5])
    if err != nil {
        return errors.Wrap(err, "entry service instance id parse error")
    }
    tc.NetworkAddress, tc.NetworkAddressID, err = decodeBase64(hh[6])
    if err != nil {
        return errors.Wrap(err, "network address parse error")
    }
    if len(hh) < 9 {
        return nil
    }
    tc.EntryEndpoint, tc.EntryEndpointID, err = decodeBase64(hh[7])
    if err != nil {
        return errors.Wrap(err, "entry endpoint parse error")
    }
    tc.ParentEndpoint, tc.ParentEndpointID, err = decodeBase64(hh[8])
    if err != nil {
        return errors.Wrap(err, "parent endpoint parse error")
    }

    fmt.Printf("\n\n===================================================Span INFO======================================\n\n: %+v\n", tc )
    fmt.Print("\n\n===================================================Span INFO==========================================\n\n" )
    return nil
}

