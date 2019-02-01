package brawndo

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	//"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

type Transport struct {
	ApiURL, Host, PublicKey, SecretKey string
	Client *http.Client
}

func (t Transport) ComputeHmac512(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha512.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func (t Transport) SignRequest( method,path,resource string, request *http.Request ) {
	var xDropoffDate = time.Now().Format("20060102T150405Z")
	var keys []string

	request.Header.Add("X-Dropoff-Date", xDropoffDate)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-agent", "brawndo-client-go")
	request.Header.Add("Host", t.Host)

	for k := range request.Header {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	var headerString, headerKeyString, authBody, bodyHash, finalStringToHash, firstKey, finalHash, authHash string

	for _,v := range keys {
		if headerString != "" {
			headerString += "\n"
			headerKeyString += ";"
		}
		headerKeyString += strings.ToLower(v)
		headerString += strings.ToLower(v)
		headerString += ":"
		headerString += request.Header.Get(v)
	}

	if headerString != "" {
		headerString += "\n"
	}

	authBody = method + "\n" + path + "\n\n" + headerString + "\n" + headerKeyString + "\n"

	bodyHash = t.ComputeHmac512(authBody, t.SecretKey)

	finalStringToHash = "HMAC-SHA512\n" + xDropoffDate + "\n" + resource + "\n" + bodyHash

	firstKey = "dropoff" + t.SecretKey
	finalHash = t.ComputeHmac512(xDropoffDate[:8], firstKey)
	finalHash = t.ComputeHmac512(resource, finalHash)
	authHash = t.ComputeHmac512(finalStringToHash, finalHash)

	var authHeader string
	authHeader = "Authorization: HMAC-SHA512 Credential=" + t.PublicKey
	authHeader += ",SignedHeaders=" + headerKeyString
	authHeader += ",Signature=" + authHash

	request.Header.Add("Authorization", authHeader)
}


func (t Transport) MakeRequest( method,path,resource,query string, body []byte ) (string, error) {
	if t.Client == nil {
		t.Client = &http.Client{}
	}

	var req *http.Request

	if body != nil {
		reqq, err := http.NewRequest(method, t.ApiURL + path + query, bytes.NewBuffer(body))

		if err != nil {
			return "", err
		}

		req = reqq
		req.Header.Set("Content-Type", "application/json")
	} else {
		reqq, err := http.NewRequest(method, t.ApiURL + path + query, nil)

		if err != nil {
			return "", err
		}

		req = reqq
	}

	t.SignRequest(method,path,resource,req)

	resp, err := t.Client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(contents), nil
}