package main

import(
        "net/http"
        "encoding/xml"
        "fmt"
        "io/ioutil"
)

type Item struct{
        Title string `xml:"title"`
        Link string `xml:"link"`
        Url string `xml:"url"`
        Name string `xml:"name"`
}
type Channel_struct struct{
        Items []Item `xml:"item"`

}
type Result struct{
        XMLName xml.Name `xml:"rss"`
        Channel Channel_struct `xml:"channel"`
}

func getRssFeed() ([]byte, error){
        //Creates a http client and request
        client := &http.Client{}
        req, err := http.NewRequest("GET", "https://ctftime.org/event/list/upcoming/rss/", nil)
        if err != nil{
                fmt.Println(err)
                return nil, err
        }
        req.Header.Add("User-Agent", "Rss-client")
        resp, err := client.Do(req)
        if err != nil{
                fmt.Println(err)
                return nil, err
        }
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil{
                fmt.Println(err)
                return nil, err
        }
        return body, nil
}
func ParseXml(data []byte) (Result, error){
        v := Result{}
        err := xml.Unmarshal(data, &v)
        if err != nil{
                fmt.Println("ERROR FROM PARSE")
                fmt.Println(err)
                return v, err
        }
        return v, nil
}
func (r *Result) iter_feed(){
        for k := range r.Channel.Items{
                fmt.Println("")
                fmt.Println(r.Channel.Items[k].Title)
                fmt.Println(r.Channel.Items[k].Link)
                fmt.Println(r.Channel.Items[k].Url)
        }
}

func main(){
    feed, err := getRssFeed()
    if err != nil{
            return
    }
    parsed_feed, err := ParseXml(feed)
    if err != nil{
            return
    }
    fmt.Printf("Items: %d\n", len(parsed_feed.Channel.Items))
    parsed_feed.iter_feed()
}

