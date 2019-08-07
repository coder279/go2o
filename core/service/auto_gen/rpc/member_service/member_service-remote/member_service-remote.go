// Autogenerated by Thrift Compiler (0.12.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
        "context"
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        "github.com/apache/thrift/lib/go/thrift"
	"go2o/core/service/auto_gen/rpc/ttype"
	"go2o/core/service/auto_gen/rpc/message_service"
        "go2o/core/service/auto_gen/rpc/member_service"
)

var _ = ttype.GoUnusedProtection__
var _ = message_service.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  Result RegisterMemberV2(string user, string pwd, i32 flag, string name, string phone, string email, string avatar,  extend)")
  fmt.Fprintln(os.Stderr, "  Result CheckLogin(string user, string pwd, bool update)")
  fmt.Fprintln(os.Stderr, "  Result CheckTradePwd(i64 id, string tradePwd)")
  fmt.Fprintln(os.Stderr, "   LevelList()")
  fmt.Fprintln(os.Stderr, "  STrustedInfo GetTrustInfo(i64 id)")
  fmt.Fprintln(os.Stderr, "  SLevel GetLevel(i32 id)")
  fmt.Fprintln(os.Stderr, "  SLevel GetLevelBySign(string sign)")
  fmt.Fprintln(os.Stderr, "  i64 SwapMemberId(ECredentials cred, string value)")
  fmt.Fprintln(os.Stderr, "  SMember GetMember(i64 id)")
  fmt.Fprintln(os.Stderr, "  SMember GetMemberByUser(string user)")
  fmt.Fprintln(os.Stderr, "  SProfile GetProfile(i64 id)")
  fmt.Fprintln(os.Stderr, "  Result Active(i64 memberId)")
  fmt.Fprintln(os.Stderr, "  Result Lock(i64 memberId, bool lock, string remark)")
  fmt.Fprintln(os.Stderr, "  Result GrantFlag(i64 memberId, i32 flag)")
  fmt.Fprintln(os.Stderr, "  SComplexMember Complex(i64 memberId)")
  fmt.Fprintln(os.Stderr, "  Result SendCode(i64 memberId, string operation, EMessageChannel msgType)")
  fmt.Fprintln(os.Stderr, "  Result CompareCode(i64 memberId, string code)")
  fmt.Fprintln(os.Stderr, "   ReceiptsCodes(i64 memberId)")
  fmt.Fprintln(os.Stderr, "  Result SaveReceiptsCode(i64 memberId, SReceiptsCode code)")
  fmt.Fprintln(os.Stderr, "   Bankcards(i64 memberId)")
  fmt.Fprintln(os.Stderr, "  Result SaveBankcard(i64 memberId, SBankcard card)")
  fmt.Fprintln(os.Stderr, "  Result CheckProfileComplete(i64 memberId)")
  fmt.Fprintln(os.Stderr, "  SMemberLevelInfo MemberLevelInfo(i64 memberId)")
  fmt.Fprintln(os.Stderr, "  Result UpdateLevel(i64 memberId, i32 level, bool review, i64 paymentOrderId)")
  fmt.Fprintln(os.Stderr, "  Result ChangePhone(i64 memberId, string phone)")
  fmt.Fprintln(os.Stderr, "  Result ChangeUsr(i64 memberId, string usr)")
  fmt.Fprintln(os.Stderr, "  Result ModifyPwd(i64 memberId, string old, string pwd)")
  fmt.Fprintln(os.Stderr, "  Result ModifyTradePwd(i64 memberId, string old, string pwd)")
  fmt.Fprintln(os.Stderr, "  Result Premium(i64 memberId, i32 v, i64 expires)")
  fmt.Fprintln(os.Stderr, "  string GetToken(i64 memberId, bool reset)")
  fmt.Fprintln(os.Stderr, "  bool CheckToken(i64 memberId, string token)")
  fmt.Fprintln(os.Stderr, "  void RemoveToken(i64 memberId)")
  fmt.Fprintln(os.Stderr, "   GetAddressList(i64 memberId)")
  fmt.Fprintln(os.Stderr, "  SAddress GetAddress(i64 memberId, i64 addrId)")
  fmt.Fprintln(os.Stderr, "  SAccount GetAccount(i64 memberId)")
  fmt.Fprintln(os.Stderr, "   InviterArray(i64 memberId, i32 depth)")
  fmt.Fprintln(os.Stderr, "  i32 InviteMembersQuantity(i64 memberId, i32 depth)")
  fmt.Fprintln(os.Stderr, "  i32 QueryInviteQuantity(i64 memberId,  data)")
  fmt.Fprintln(os.Stderr, "   QueryInviteArray(i64 memberId,  data)")
  fmt.Fprintln(os.Stderr, "  Result AccountCharge(i64 memberId, i32 account, string title, i32 amount, string outerNo, string remark)")
  fmt.Fprintln(os.Stderr, "  Result AccountConsume(i64 memberId, i32 account, string title, i32 amount, string outerNo, string remark)")
  fmt.Fprintln(os.Stderr, "  Result AccountDiscount(i64 memberId, i32 account, string title, i32 amount, string outerNo, string remark)")
  fmt.Fprintln(os.Stderr, "  Result AccountRefund(i64 memberId, i32 account, string title, i32 amount, string outerNo, string remark)")
  fmt.Fprintln(os.Stderr, "  Result AccountAdjust(i64 memberId, i32 account, i32 value, i64 relateUser, string remark)")
  fmt.Fprintln(os.Stderr, "  Result B4EAuth(i64 memberId, string action,  data)")
  fmt.Fprintln(os.Stderr, "  SPagingResult PagingAccountLog(i64 memberId, i32 accountType, SPagingParams params)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
  var m map[string]string = h
  return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
  parts := strings.Split(value, ": ")
  if len(parts) != 2 {
    return fmt.Errorf("header should be of format 'Key: Value'")
  }
  h[parts[0]] = parts[1]
  return nil
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  headers := make(httpHeaders)
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
    if len(headers) > 0 {
      httptrans := trans.(*thrift.THttpClient)
      for key, value := range headers {
        httptrans.SetHeader(key, value)
      }
    }
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := member_service.NewMemberServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "RegisterMemberV2":
    if flag.NArg() - 1 != 8 {
      fmt.Fprintln(os.Stderr, "RegisterMemberV2 requires 8 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    tmp2, err110 := (strconv.Atoi(flag.Arg(3)))
    if err110 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    argvalue3 := flag.Arg(4)
    value3 := argvalue3
    argvalue4 := flag.Arg(5)
    value4 := argvalue4
    argvalue5 := flag.Arg(6)
    value5 := argvalue5
    argvalue6 := flag.Arg(7)
    value6 := argvalue6
    arg115 := flag.Arg(8)
    mbTrans116 := thrift.NewTMemoryBufferLen(len(arg115))
    defer mbTrans116.Close()
    _, err117 := mbTrans116.WriteString(arg115)
    if err117 != nil { 
      Usage()
      return
    }
    factory118 := thrift.NewTJSONProtocolFactory()
    jsProt119 := factory118.GetProtocol(mbTrans116)
    containerStruct7 := member_service.NewMemberServiceRegisterMemberV2Args()
    err120 := containerStruct7.ReadField8(jsProt119)
    if err120 != nil {
      Usage()
      return
    }
    argvalue7 := containerStruct7.Extend
    value7 := argvalue7
    fmt.Print(client.RegisterMemberV2(context.Background(), value0, value1, value2, value3, value4, value5, value6, value7))
    fmt.Print("\n")
    break
  case "CheckLogin":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "CheckLogin requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3) == "true"
    value2 := argvalue2
    fmt.Print(client.CheckLogin(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "CheckTradePwd":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "CheckTradePwd requires 2 args")
      flag.Usage()
    }
    argvalue0, err124 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err124 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.CheckTradePwd(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "LevelList":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "LevelList requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.LevelList(context.Background()))
    fmt.Print("\n")
    break
  case "GetTrustInfo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTrustInfo requires 1 args")
      flag.Usage()
    }
    argvalue0, err126 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err126 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetTrustInfo(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetLevel":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetLevel requires 1 args")
      flag.Usage()
    }
    tmp0, err127 := (strconv.Atoi(flag.Arg(1)))
    if err127 != nil {
      Usage()
      return
    }
    argvalue0 := int32(tmp0)
    value0 := argvalue0
    fmt.Print(client.GetLevel(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetLevelBySign":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetLevelBySign requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GetLevelBySign(context.Background(), value0))
    fmt.Print("\n")
    break
  case "SwapMemberId":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SwapMemberId requires 2 args")
      flag.Usage()
    }
    tmp0, err := (strconv.Atoi(flag.Arg(1)))
    if err != nil {
      Usage()
     return
    }
    argvalue0 := member_service.ECredentials(tmp0)
    value0 := member_service.ECredentials(argvalue0)
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.SwapMemberId(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "GetMember":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetMember requires 1 args")
      flag.Usage()
    }
    argvalue0, err130 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err130 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetMember(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetMemberByUser":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetMemberByUser requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GetMemberByUser(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetProfile":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetProfile requires 1 args")
      flag.Usage()
    }
    argvalue0, err132 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err132 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetProfile(context.Background(), value0))
    fmt.Print("\n")
    break
  case "Active":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Active requires 1 args")
      flag.Usage()
    }
    argvalue0, err133 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err133 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Active(context.Background(), value0))
    fmt.Print("\n")
    break
  case "Lock":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "Lock requires 3 args")
      flag.Usage()
    }
    argvalue0, err134 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err134 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2) == "true"
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.Lock(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "GrantFlag":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "GrantFlag requires 2 args")
      flag.Usage()
    }
    argvalue0, err137 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err137 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err138 := (strconv.Atoi(flag.Arg(2)))
    if err138 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.GrantFlag(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "Complex":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Complex requires 1 args")
      flag.Usage()
    }
    argvalue0, err139 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err139 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Complex(context.Background(), value0))
    fmt.Print("\n")
    break
  case "SendCode":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "SendCode requires 3 args")
      flag.Usage()
    }
    argvalue0, err140 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err140 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    tmp2, err := (strconv.Atoi(flag.Arg(3)))
    if err != nil {
      Usage()
     return
    }
    argvalue2 := member_service.EMessageChannel(tmp2)
    value2 := argvalue2
    fmt.Print(client.SendCode(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "CompareCode":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "CompareCode requires 2 args")
      flag.Usage()
    }
    argvalue0, err142 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err142 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.CompareCode(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "ReceiptsCodes":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ReceiptsCodes requires 1 args")
      flag.Usage()
    }
    argvalue0, err144 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err144 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.ReceiptsCodes(context.Background(), value0))
    fmt.Print("\n")
    break
  case "SaveReceiptsCode":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SaveReceiptsCode requires 2 args")
      flag.Usage()
    }
    argvalue0, err145 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err145 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    arg146 := flag.Arg(2)
    mbTrans147 := thrift.NewTMemoryBufferLen(len(arg146))
    defer mbTrans147.Close()
    _, err148 := mbTrans147.WriteString(arg146)
    if err148 != nil {
      Usage()
      return
    }
    factory149 := thrift.NewTJSONProtocolFactory()
    jsProt150 := factory149.GetProtocol(mbTrans147)
    argvalue1 := member_service.NewSReceiptsCode()
    err151 := argvalue1.Read(jsProt150)
    if err151 != nil {
      Usage()
      return
    }
    value1 := member_service.SReceiptsCode(argvalue1)
    fmt.Print(client.SaveReceiptsCode(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "Bankcards":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Bankcards requires 1 args")
      flag.Usage()
    }
    argvalue0, err152 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err152 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.Bankcards(context.Background(), value0))
    fmt.Print("\n")
    break
  case "SaveBankcard":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SaveBankcard requires 2 args")
      flag.Usage()
    }
    argvalue0, err153 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err153 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    arg154 := flag.Arg(2)
    mbTrans155 := thrift.NewTMemoryBufferLen(len(arg154))
    defer mbTrans155.Close()
    _, err156 := mbTrans155.WriteString(arg154)
    if err156 != nil {
      Usage()
      return
    }
    factory157 := thrift.NewTJSONProtocolFactory()
    jsProt158 := factory157.GetProtocol(mbTrans155)
    argvalue1 := member_service.NewSBankcard()
    err159 := argvalue1.Read(jsProt158)
    if err159 != nil {
      Usage()
      return
    }
    value1 := member_service.SBankcard(argvalue1)
    fmt.Print(client.SaveBankcard(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "CheckProfileComplete":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CheckProfileComplete requires 1 args")
      flag.Usage()
    }
    argvalue0, err160 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err160 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CheckProfileComplete(context.Background(), value0))
    fmt.Print("\n")
    break
  case "MemberLevelInfo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "MemberLevelInfo requires 1 args")
      flag.Usage()
    }
    argvalue0, err161 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err161 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.MemberLevelInfo(context.Background(), value0))
    fmt.Print("\n")
    break
  case "UpdateLevel":
    if flag.NArg() - 1 != 4 {
      fmt.Fprintln(os.Stderr, "UpdateLevel requires 4 args")
      flag.Usage()
    }
    argvalue0, err162 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err162 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err163 := (strconv.Atoi(flag.Arg(2)))
    if err163 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    argvalue2 := flag.Arg(3) == "true"
    value2 := argvalue2
    argvalue3, err165 := (strconv.ParseInt(flag.Arg(4), 10, 64))
    if err165 != nil {
      Usage()
      return
    }
    value3 := argvalue3
    fmt.Print(client.UpdateLevel(context.Background(), value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "ChangePhone":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "ChangePhone requires 2 args")
      flag.Usage()
    }
    argvalue0, err166 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err166 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.ChangePhone(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "ChangeUsr":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "ChangeUsr requires 2 args")
      flag.Usage()
    }
    argvalue0, err168 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err168 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.ChangeUsr(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "ModifyPwd":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "ModifyPwd requires 3 args")
      flag.Usage()
    }
    argvalue0, err170 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err170 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.ModifyPwd(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "ModifyTradePwd":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "ModifyTradePwd requires 3 args")
      flag.Usage()
    }
    argvalue0, err173 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err173 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.ModifyTradePwd(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "Premium":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "Premium requires 3 args")
      flag.Usage()
    }
    argvalue0, err176 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err176 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err177 := (strconv.Atoi(flag.Arg(2)))
    if err177 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    argvalue2, err178 := (strconv.ParseInt(flag.Arg(3), 10, 64))
    if err178 != nil {
      Usage()
      return
    }
    value2 := argvalue2
    fmt.Print(client.Premium(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "GetToken":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "GetToken requires 2 args")
      flag.Usage()
    }
    argvalue0, err179 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err179 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2) == "true"
    value1 := argvalue1
    fmt.Print(client.GetToken(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "CheckToken":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "CheckToken requires 2 args")
      flag.Usage()
    }
    argvalue0, err181 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err181 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.CheckToken(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "RemoveToken":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "RemoveToken requires 1 args")
      flag.Usage()
    }
    argvalue0, err183 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err183 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.RemoveToken(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetAddressList":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetAddressList requires 1 args")
      flag.Usage()
    }
    argvalue0, err184 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err184 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetAddressList(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetAddress":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "GetAddress requires 2 args")
      flag.Usage()
    }
    argvalue0, err185 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err185 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1, err186 := (strconv.ParseInt(flag.Arg(2), 10, 64))
    if err186 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.GetAddress(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "GetAccount":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetAccount requires 1 args")
      flag.Usage()
    }
    argvalue0, err187 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err187 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetAccount(context.Background(), value0))
    fmt.Print("\n")
    break
  case "InviterArray":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "InviterArray requires 2 args")
      flag.Usage()
    }
    argvalue0, err188 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err188 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err189 := (strconv.Atoi(flag.Arg(2)))
    if err189 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.InviterArray(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "InviteMembersQuantity":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "InviteMembersQuantity requires 2 args")
      flag.Usage()
    }
    argvalue0, err190 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err190 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err191 := (strconv.Atoi(flag.Arg(2)))
    if err191 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.InviteMembersQuantity(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "QueryInviteQuantity":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "QueryInviteQuantity requires 2 args")
      flag.Usage()
    }
    argvalue0, err192 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err192 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    arg193 := flag.Arg(2)
    mbTrans194 := thrift.NewTMemoryBufferLen(len(arg193))
    defer mbTrans194.Close()
    _, err195 := mbTrans194.WriteString(arg193)
    if err195 != nil { 
      Usage()
      return
    }
    factory196 := thrift.NewTJSONProtocolFactory()
    jsProt197 := factory196.GetProtocol(mbTrans194)
    containerStruct1 := member_service.NewMemberServiceQueryInviteQuantityArgs()
    err198 := containerStruct1.ReadField2(jsProt197)
    if err198 != nil {
      Usage()
      return
    }
    argvalue1 := containerStruct1.Data
    value1 := argvalue1
    fmt.Print(client.QueryInviteQuantity(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "QueryInviteArray":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "QueryInviteArray requires 2 args")
      flag.Usage()
    }
    argvalue0, err199 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err199 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    arg200 := flag.Arg(2)
    mbTrans201 := thrift.NewTMemoryBufferLen(len(arg200))
    defer mbTrans201.Close()
    _, err202 := mbTrans201.WriteString(arg200)
    if err202 != nil { 
      Usage()
      return
    }
    factory203 := thrift.NewTJSONProtocolFactory()
    jsProt204 := factory203.GetProtocol(mbTrans201)
    containerStruct1 := member_service.NewMemberServiceQueryInviteArrayArgs()
    err205 := containerStruct1.ReadField2(jsProt204)
    if err205 != nil {
      Usage()
      return
    }
    argvalue1 := containerStruct1.Data
    value1 := argvalue1
    fmt.Print(client.QueryInviteArray(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "AccountCharge":
    if flag.NArg() - 1 != 6 {
      fmt.Fprintln(os.Stderr, "AccountCharge requires 6 args")
      flag.Usage()
    }
    argvalue0, err206 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err206 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err207 := (strconv.Atoi(flag.Arg(2)))
    if err207 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    tmp3, err209 := (strconv.Atoi(flag.Arg(4)))
    if err209 != nil {
      Usage()
      return
    }
    argvalue3 := int32(tmp3)
    value3 := argvalue3
    argvalue4 := flag.Arg(5)
    value4 := argvalue4
    argvalue5 := flag.Arg(6)
    value5 := argvalue5
    fmt.Print(client.AccountCharge(context.Background(), value0, value1, value2, value3, value4, value5))
    fmt.Print("\n")
    break
  case "AccountConsume":
    if flag.NArg() - 1 != 6 {
      fmt.Fprintln(os.Stderr, "AccountConsume requires 6 args")
      flag.Usage()
    }
    argvalue0, err212 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err212 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err213 := (strconv.Atoi(flag.Arg(2)))
    if err213 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    tmp3, err215 := (strconv.Atoi(flag.Arg(4)))
    if err215 != nil {
      Usage()
      return
    }
    argvalue3 := int32(tmp3)
    value3 := argvalue3
    argvalue4 := flag.Arg(5)
    value4 := argvalue4
    argvalue5 := flag.Arg(6)
    value5 := argvalue5
    fmt.Print(client.AccountConsume(context.Background(), value0, value1, value2, value3, value4, value5))
    fmt.Print("\n")
    break
  case "AccountDiscount":
    if flag.NArg() - 1 != 6 {
      fmt.Fprintln(os.Stderr, "AccountDiscount requires 6 args")
      flag.Usage()
    }
    argvalue0, err218 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err218 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err219 := (strconv.Atoi(flag.Arg(2)))
    if err219 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    tmp3, err221 := (strconv.Atoi(flag.Arg(4)))
    if err221 != nil {
      Usage()
      return
    }
    argvalue3 := int32(tmp3)
    value3 := argvalue3
    argvalue4 := flag.Arg(5)
    value4 := argvalue4
    argvalue5 := flag.Arg(6)
    value5 := argvalue5
    fmt.Print(client.AccountDiscount(context.Background(), value0, value1, value2, value3, value4, value5))
    fmt.Print("\n")
    break
  case "AccountRefund":
    if flag.NArg() - 1 != 6 {
      fmt.Fprintln(os.Stderr, "AccountRefund requires 6 args")
      flag.Usage()
    }
    argvalue0, err224 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err224 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err225 := (strconv.Atoi(flag.Arg(2)))
    if err225 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    tmp3, err227 := (strconv.Atoi(flag.Arg(4)))
    if err227 != nil {
      Usage()
      return
    }
    argvalue3 := int32(tmp3)
    value3 := argvalue3
    argvalue4 := flag.Arg(5)
    value4 := argvalue4
    argvalue5 := flag.Arg(6)
    value5 := argvalue5
    fmt.Print(client.AccountRefund(context.Background(), value0, value1, value2, value3, value4, value5))
    fmt.Print("\n")
    break
  case "AccountAdjust":
    if flag.NArg() - 1 != 5 {
      fmt.Fprintln(os.Stderr, "AccountAdjust requires 5 args")
      flag.Usage()
    }
    argvalue0, err230 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err230 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err231 := (strconv.Atoi(flag.Arg(2)))
    if err231 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    tmp2, err232 := (strconv.Atoi(flag.Arg(3)))
    if err232 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    argvalue3, err233 := (strconv.ParseInt(flag.Arg(4), 10, 64))
    if err233 != nil {
      Usage()
      return
    }
    value3 := argvalue3
    argvalue4 := flag.Arg(5)
    value4 := argvalue4
    fmt.Print(client.AccountAdjust(context.Background(), value0, value1, value2, value3, value4))
    fmt.Print("\n")
    break
  case "B4EAuth":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "B4EAuth requires 3 args")
      flag.Usage()
    }
    argvalue0, err235 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err235 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    arg237 := flag.Arg(3)
    mbTrans238 := thrift.NewTMemoryBufferLen(len(arg237))
    defer mbTrans238.Close()
    _, err239 := mbTrans238.WriteString(arg237)
    if err239 != nil { 
      Usage()
      return
    }
    factory240 := thrift.NewTJSONProtocolFactory()
    jsProt241 := factory240.GetProtocol(mbTrans238)
    containerStruct2 := member_service.NewMemberServiceB4EAuthArgs()
    err242 := containerStruct2.ReadField3(jsProt241)
    if err242 != nil {
      Usage()
      return
    }
    argvalue2 := containerStruct2.Data
    value2 := argvalue2
    fmt.Print(client.B4EAuth(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "PagingAccountLog":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "PagingAccountLog requires 3 args")
      flag.Usage()
    }
    argvalue0, err243 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err243 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err244 := (strconv.Atoi(flag.Arg(2)))
    if err244 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    arg245 := flag.Arg(3)
    mbTrans246 := thrift.NewTMemoryBufferLen(len(arg245))
    defer mbTrans246.Close()
    _, err247 := mbTrans246.WriteString(arg245)
    if err247 != nil {
      Usage()
      return
    }
    factory248 := thrift.NewTJSONProtocolFactory()
    jsProt249 := factory248.GetProtocol(mbTrans246)
    argvalue2 := ttype.NewSPagingParams()
    err250 := argvalue2.Read(jsProt249)
    if err250 != nil {
      Usage()
      return
    }
    value2 := argvalue2
    fmt.Print(client.PagingAccountLog(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
