// Autogenerated by Thrift Compiler (0.13.0)
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
	"github.com/diffeo/goevernote/edam"
)

var _ = edam.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  bool checkVersion(string clientName, i16 edamVersionMajor, i16 edamVersionMinor)")
  fmt.Fprintln(os.Stderr, "  BootstrapInfo getBootstrapInfo(string locale)")
  fmt.Fprintln(os.Stderr, "  AuthenticationResult authenticateLongSession(string username, string password, string consumerKey, string consumerSecret, string deviceIdentifier, string deviceDescription, bool supportsTwoFactor)")
  fmt.Fprintln(os.Stderr, "  AuthenticationResult completeTwoFactorAuthentication(string authenticationToken, string oneTimeCode, string deviceIdentifier, string deviceDescription)")
  fmt.Fprintln(os.Stderr, "  void revokeLongSession(string authenticationToken)")
  fmt.Fprintln(os.Stderr, "  AuthenticationResult authenticateToBusiness(string authenticationToken)")
  fmt.Fprintln(os.Stderr, "  User getUser(string authenticationToken)")
  fmt.Fprintln(os.Stderr, "  PublicUserInfo getPublicUserInfo(string username)")
  fmt.Fprintln(os.Stderr, "  UserUrls getUserUrls(string authenticationToken)")
  fmt.Fprintln(os.Stderr, "  void inviteToBusiness(string authenticationToken, string emailAddress)")
  fmt.Fprintln(os.Stderr, "  void removeFromBusiness(string authenticationToken, string emailAddress)")
  fmt.Fprintln(os.Stderr, "  void updateBusinessUserIdentifier(string authenticationToken, string oldEmailAddress, string newEmailAddress)")
  fmt.Fprintln(os.Stderr, "   listBusinessUsers(string authenticationToken)")
  fmt.Fprintln(os.Stderr, "   listBusinessInvitations(string authenticationToken, bool includeRequestedInvitations)")
  fmt.Fprintln(os.Stderr, "  AccountLimits getAccountLimits(ServiceLevel serviceLevel)")
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
  client := edam.NewUserStoreClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "checkVersion":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "CheckVersion requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    tmp1, err36 := (strconv.Atoi(flag.Arg(2)))
    if err36 != nil {
      Usage()
      return
    }
    argvalue1 := int16(tmp1)
    value1 := argvalue1
    tmp2, err37 := (strconv.Atoi(flag.Arg(3)))
    if err37 != nil {
      Usage()
      return
    }
    argvalue2 := int16(tmp2)
    value2 := argvalue2
    fmt.Print(client.CheckVersion(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "getBootstrapInfo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetBootstrapInfo requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GetBootstrapInfo(context.Background(), value0))
    fmt.Print("\n")
    break
  case "authenticateLongSession":
    if flag.NArg() - 1 != 7 {
      fmt.Fprintln(os.Stderr, "AuthenticateLongSession requires 7 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    argvalue3 := flag.Arg(4)
    value3 := argvalue3
    argvalue4 := flag.Arg(5)
    value4 := argvalue4
    argvalue5 := flag.Arg(6)
    value5 := argvalue5
    argvalue6 := flag.Arg(7) == "true"
    value6 := argvalue6
    fmt.Print(client.AuthenticateLongSession(context.Background(), value0, value1, value2, value3, value4, value5, value6))
    fmt.Print("\n")
    break
  case "completeTwoFactorAuthentication":
    if flag.NArg() - 1 != 4 {
      fmt.Fprintln(os.Stderr, "CompleteTwoFactorAuthentication requires 4 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    argvalue3 := flag.Arg(4)
    value3 := argvalue3
    fmt.Print(client.CompleteTwoFactorAuthentication(context.Background(), value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "revokeLongSession":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "RevokeLongSession requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.RevokeLongSession(context.Background(), value0))
    fmt.Print("\n")
    break
  case "authenticateToBusiness":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "AuthenticateToBusiness requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.AuthenticateToBusiness(context.Background(), value0))
    fmt.Print("\n")
    break
  case "getUser":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetUser requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GetUser(context.Background(), value0))
    fmt.Print("\n")
    break
  case "getPublicUserInfo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetPublicUserInfo requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GetPublicUserInfo(context.Background(), value0))
    fmt.Print("\n")
    break
  case "getUserUrls":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetUserUrls requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GetUserUrls(context.Background(), value0))
    fmt.Print("\n")
    break
  case "inviteToBusiness":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "InviteToBusiness requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.InviteToBusiness(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "removeFromBusiness":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "RemoveFromBusiness requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.RemoveFromBusiness(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "updateBusinessUserIdentifier":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "UpdateBusinessUserIdentifier requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.UpdateBusinessUserIdentifier(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "listBusinessUsers":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ListBusinessUsers requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.ListBusinessUsers(context.Background(), value0))
    fmt.Print("\n")
    break
  case "listBusinessInvitations":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "ListBusinessInvitations requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2) == "true"
    value1 := argvalue1
    fmt.Print(client.ListBusinessInvitations(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "getAccountLimits":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetAccountLimits requires 1 args")
      flag.Usage()
    }
    tmp0, err := (strconv.Atoi(flag.Arg(1)))
    if err != nil {
      Usage()
     return
    }
    argvalue0 := edam.ServiceLevel(tmp0)
    value0 := argvalue0
    fmt.Print(client.GetAccountLimits(context.Background(), value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
