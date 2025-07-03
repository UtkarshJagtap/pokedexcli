package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)
func SendRequest(ty ,url string) ( []byte, error){

  res, err := http.Get(url)
  if err != nil{
    return nil, err
  }
  defer res.Body.Close()

  var body []byte

  if res.StatusCode > 199 && res.StatusCode < 300{
    body, err = io.ReadAll(res.Body)
    if err != nil{
      return nil, err
    }
  }

  if res.StatusCode > 399 && res.StatusCode < 500{
    return nil, fmt.Errorf("bad request or resource not found %v",res.Status)
  }

  return body, nil
}
