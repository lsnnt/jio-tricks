package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"flag"
)

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
	req.Header.Set("accesstoken", "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImF1dGhUb2tlbklkIjoiMzI2YjY3YTQtNjlhNC00ZmJmLTgwNmMtYjM1MTA5ZWE4YTRhIiwidXNlcklkIjoiODE1ZDFiODQtM2M2ZC00YjI5LTk4NDktNDBmZDg0MmI4YjNlIiwidXNlclR5cGUiOiJOT05KSU8iLCJvcyI6IndlYiIsImRldmljZVR5cGUiOiJwYyIsImFjY2Vzc0xldmVsIjoiOSIsImRldmljZUlkIjoiMzI0YzVjOGEtMjJjZC00NmY3LWEwMjItYzkzOTJmMGQ0NzQ5IiwiZXh0cmEiOiJ7XCJudW1iZXJcIjpcIk9sUDBaclRhOHJUZ3U4NXh4REFDVXI3OHBqNGJoWnU0UjliQUxqWW54Y1Q2QlBDSFdzYlRQMkk9XCIsXCJhZHNcIjpcInllc1wiLFwicGxhbmRldGFpbHNcIjp7XCJhZHNcIjpcInllc1wiLFwiUGFja2FnZUluZm9cIjpbXX0sXCJqVG9rZW5cIjpcIlwiLFwidXNlckRldGFpbHNcIjpcIm1DUUdGUktJZ0lVMEppTXd5aURsODFMMkNndTlsTGc1QWRhU2Yrbjd6cVlzQ2VrV3NZL3hMeFVGazNQWHg2WlE0WjMzczBVdWxsZ0Y5Nm1GYzFNMURpcWsvc1VSdnUyZStlYWxYdlQ2Syt6NXRqU3BERWs0TjBvUmszelovS29BOWpKb1hNZ0dWZ1JHcFloYVdoL0labXZxRlRQRnZJd1lvWnZ5TE9HeFl3PT1cIn0iLCJzdWJzY3JpYmVySWQiOiIiLCJhcHBOYW1lIjoiUkpJTF9KaW9DaW5lbWEiLCJkZWdyYWRlZCI6ImZhbHNlIiwiYWRzIjoieWVzIiwicHJvZmlsZUlkIjoiNmYyOTU4NDktMWY0YS00ZjVjLTkxOTMtMWZkODE3ZmRhZDU2IiwiYWRJZCI6IjMyNGM1YzhhLTIyY2QtNDZmNy1hMDIyLWM5MzkyZjBkNDc0OSIsImFkc0NvbmZpZyI6eyJpbnN0cmVhbUFkcyI6eyJsaXZlIjp7ImVuYWJsZWQiOnRydWV9LCJ2b2QiOnsiZW5hYmxlZCI6dHJ1ZX19fSwiZXhwZXJpbWVudEtleSI6eyJjb25maWdLZXkiOiI2ZjI5NTg0OS0xZjRhLTRmNWMtOTE5My0xZmQ4MTdmZGFkNTYiLCJncm91cElkIjozODE3fSwicHJvZmlsZURldGFpbHMiOnsicHJvZmlsZVR5cGUiOiJjaGlsZCIsImNvbnRlbnRBZ2VSYXRpbmciOiJVL0EgNysifSwidmVyc2lvbiI6MjAyNDAzMDQwfSwiZXhwIjoxNzE3MzEwNzc0LCJpYXQiOjE3MTczMDM1NzR9.SDxPBDPTLLno9JkrrfbeWeMuHO8nn2k9OSpv5bISVKg10Oy3vS3FIKgIqNPHxe5vMM3ZdqMu8jHQ291Ca9IGLw")
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
}
