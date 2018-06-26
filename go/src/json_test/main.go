package main

import (
  "fmt"
  "encoding/json"
)

type ref struct {
  Commit string `json:commit`
  PreviousCommit string `json:previousCommit`
  Repository string `json:repository`
}

type body struct {
  Version string `json:version`
  Refs   []*ref `json:refs`
  Projects []string `json:projects`
}
func main() {
  s := `{"version": "c8cb35287","refs": [{"commit":"c8cb35287","previousCommit":"5d646af2","repository":"Repo/repo"}],"projects":["core"]}`
  t := []byte(s)
  var b body
  err := json.Unmarshal(t, &b)

  if err != nil {
    fmt.Println(err)
  }
  fmt.Printf("%s \n", b.Version)
  for _,i := range b.Refs {
    fmt.Printf("Commit :%s , Previous commit : %s", i.Commit, i.PreviousCommit)
  }
}
