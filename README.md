# go-rest-api

```Go
func main() {
  var err error

  config, err := config.BuildConfig(".env")
  if err != nil {
    panic(err)
  }

  client, err := pexip.NewForConfig(config)
  if err != nil {
    panic(err)
  }

  clientset := client.ManagementV1()

  err = clientset.Command().Participant().Dial()
  if err != nil {
    panic(err)
  }
}
```
