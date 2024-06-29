package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"flag"
	"encoding/json"
)
type UserData struct {
    AuthToken string `json:"authToken"`
    // Add other fields if needed
}
func getguestauthtoken() string{
	client := &http.Client{}
	var data = strings.NewReader(`{"appName":"RJIL_JioCinema","deviceType":"pc","os":"web","deviceId":"aabf1b9a-3ae8-43a0-b0f0-b3062ea4bea4","freshLaunch":true,"adId":"aabf1b9a-3ae8-43a0-b0f0-b3062ea4bea4","appVersion":"24.06.25.1-d22894a9"}`)
	req, err := http.NewRequest("POST", "https://auth-jiocinema.voot.com/tokenservice/apis/v4/guest", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "en-GB,en;q=0.8")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("origin", "https://www.jiocinema.com")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("referer", "https://www.jiocinema.com/")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="8", "Chromium";v="126", "Brave";v="126"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Linux"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", bodyText)
	var dataa UserData
	if err := json.Unmarshal([]byte(bodyText), &dataa); err != nil {
        log.Fatal(err)
    }

    // Access and print the authToken field
    // fmt.Println("AuthToken:", dataa.AuthToken)
	return dataa.AuthToken
}
func main() {
	client := &http.Client{}
	var id string 
	flag.StringVar(&id, "id", "", "ID of the video")
	flag.Parse()
	if id == "" {
		log.Fatal("ID is required to fetch the video pass it by using --id flag") 
	}
	var url string
	url = fmt.Sprintf("https://apis-jiovoot.voot.com/playbackjv/v5/%s", id)

	var data = strings.NewReader(`{"4k":false,"ageGroup":"18+","appVersion":"3.4.0","bitrateProfile":"xhdpi","capability":{"drmCapability":{"aesSupport":"yes","fairPlayDrmSupport":"yes","playreadyDrmSupport":"none","widevineDRMSupport":"yes"},"frameRateCapability":[{"frameRateSupport":"30fps","videoQuality":"1440p"}]},"continueWatchingRequired":true,"dolby":false,"downloadRequest":false,"hevc":false,"kidsSafe":false,"manufacturer":"Linux","model":"Linux","multiAudioRequired":true,"osVersion":"x86_64","parentalPinValid":true,"x-apisignatures":"o668nxgzwff"}`)
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "en-GB,en;q=0.7")
	req.Header.Set("accesstoken", getguestauthtoken()) // populate authtoken here
	req.Header.Set("appname", "RJIL_JioCinema")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("deviceid", "324c5c8a-22cd-46f7-a022-c9392f0d4749")
	req.Header.Set("origin", "https://www.jiocinema.com")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("referer", "https://www.jiocinema.com/")
	req.Header.Set("sec-ch-ua", `"Brave";v="125", "Chromium";v="125", "Not.A/Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Linux"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("uniqueid", "815d1b84-3c6d-4b29-9849-40fd842b8b3e")
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")
	req.Header.Set("versioncode", "2405231")
	req.Header.Set("x-apisignatures", "o668nxgzwff")
	req.Header.Set("x-platform", "androidweb")
	req.Header.Set("x-platform-token", "web")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	dashurl := strings.Split(string(bodyText), `,{"macros":[],"url":"`)[1]
	dashurl = strings.Split(dashurl, `","`)[0]
	fmt.Println(dashurl)
	fmt.Println("\n\n use the above url in vlc media player and enjoy watching after FUP data limit ehausted")
}
