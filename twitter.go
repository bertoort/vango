package main

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/subosito/gotenv"
)

var api *anaconda.TwitterApi

func init() {
	gotenv.Load()
	anaconda.SetConsumerKey(os.Getenv("COSTUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("COSTUMER_SECRET"))
	api = anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
}

func postMedia(newImage image.Image) (media anaconda.Media, err error) {
	buf := new(bytes.Buffer)
	err = png.Encode(buf, newImage)
	if err != nil {
		return
	}
	base64ImageString := base64.StdEncoding.EncodeToString(buf.Bytes())
	media, err = api.UploadMedia(base64ImageString)
	if err != nil {
		return
	}
	return
}

func tweetMedia(media anaconda.Media) error {
	status := "Van Go Art: https://github.com/berto/vango"
	v := url.Values{}
	v.Set("media_ids", media.MediaIDString)
	_, err := api.PostTweet(status, v)
	if err != nil {
		return err
	}
	return nil
}
