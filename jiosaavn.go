package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func getsecret(url string) string {
        client := &http.Client{}
        //getting input from user
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
                log.Fatal(err)
        }
        req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
        req.Header.Set("Accept-Language", "en-GB,en;q=0.9")
        req.Header.Set("Cache-Control", "no-cache")
        req.Header.Set("Connection", "keep-alive")
        req.Header.Set("Cookie", "DL=english; B=f160d71be762e55c0ec9e353a01cb136; CT=ODc2MTIzNzk%3D; L=hindi; geo=2409%3A40c1%3A3b%3Ae9fa%3A79b4%3A7260%3Afd67%3A3f0b%2CIN%2CGujarat%2CAhmedabad%2C382350; mm_latlong=23.0276%2C72.5871; CH=G03%2CA07%2CO00%2CL03")
        req.Header.Set("Pragma", "no-cache")
        req.Header.Set("Sec-Fetch-Dest", "document")
        req.Header.Set("Sec-Fetch-Mode", "navigate")
        req.Header.Set("Sec-Fetch-Site", "none")
        req.Header.Set("Sec-Fetch-User", "?1")
        req.Header.Set("Sec-GPC", "1")
        req.Header.Set("Upgrade-Insecure-Requests", "1")
        req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")
        req.Header.Set("sec-ch-ua", `"Brave";v="125", "Chromium";v="125", "Not.A/Brand";v="24"`)
        req.Header.Set("sec-ch-ua-mobile", "?0")
        req.Header.Set("sec-ch-ua-platform", `"Linux"`)
        resp, err := client.Do(req)
        if err != nil {
                log.Fatal(err)
        }
        defer resp.Body.Close()
        bodyText, err := io.ReadAll(resp.Body)
        if err != nil {
                log.Fatal(err)
        }
        //https://www.jiosaavn.com/song/tum-ko-dekha-to-yeh-khayal-aaya/Iy4ueAQBX1k
        //https://www.jiosaavn.com/song/tum-ko-pata-hai/QAwaYT9jWFc
        splitted_body_text := strings.Split(string(bodyText),`"encrypted_media_url":"`)
        // splitted_url := strings.Split(splitted_body_text[0],"\"")
        idWithTrailingQuote := splitted_body_text[2]
    id := strings.Split(idWithTrailingQuote,"\"")[0]
    id = strings.Replace(id, "\\u002F", "/", -1)
        return id
}
func urlencode(ids string) string {
        encodedID := url.QueryEscape(ids)
    return encodedID
}
func getprecdnurl(secretencoded string) string {
        client := &http.Client{}
        url := fmt.Sprintf("https://www.jiosaavn.com/api.php?__call=song.generateAuthToken&url=%s&bitrate=128&api_version=4&_format=json&ctx=wap6dot0&_marker=0", secretencoded)
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
                log.Fatal(err)
        }
        req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
        req.Header.Set("Accept-Language", "en-GB,en;q=0.9")
        req.Header.Set("Cache-Control", "no-cache")
        req.Header.Set("Connection", "keep-alive")
        req.Header.Set("Cookie", "DL=english; B=f160d71be762e55c0ec9e353a01cb136; CT=ODc2MTIzNzk%3D; L=hindi; CH=G03%2CA07%2CO00%2CL03; geo=2401%3A4900%3A1956%3A8ca0%3A3b31%3A641a%3A5bb5%3A7cc0%2CIN%2CGujarat%2CSurat%2C394210; mm_latlong=21.1888%2C72.8293")
        req.Header.Set("Pragma", "no-cache")
        req.Header.Set("Sec-Fetch-Dest", "document")
        req.Header.Set("Sec-Fetch-Mode", "navigate")
        req.Header.Set("Sec-Fetch-Site", "none")
        req.Header.Set("Sec-Fetch-User", "?1")
        req.Header.Set("Sec-GPC", "1")
        req.Header.Set("Upgrade-Insecure-Requests", "1")
        req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")
        req.Header.Set("sec-ch-ua", `"Brave";v="125", "Chromium";v="125", "Not.A/Brand";v="24"`)
        req.Header.Set("sec-ch-ua-mobile", "?0")
        req.Header.Set("sec-ch-ua-platform", `"Linux"`)
        resp, err := client.Do(req)
        if err != nil {
                log.Fatal(err)
        }
        defer resp.Body.Close()
        bodyText, err := io.ReadAll(resp.Body)
        if err != nil {
                log.Fatal(err)
        }
        splitted_body_text := strings.Split(string(bodyText),`{"auth_url":"`)
        // splitted_url := strings.Split(splitted_body_text[0],"\"")
        idWithTrailingQuote := splitted_body_text[1]
    id := strings.Split(idWithTrailingQuote,"\"")[0]
    id = strings.Replace(id, "\\/", "/", -1)
        return id
}
func main() {
        var url string
		flag.StringVar(&url, "url", "", "URL to be encoded")
		flag.Parse()

		// Check if the URL flag is provided
		if url == "" {
			fmt.Println("Please provide a URL using the --url flag.")
			return
		}

		// Print the encoded URL
        fmt.Println(getprecdnurl(urlencode(getsecret(url))))
}
