# image service

A simple service to upload and retreive images using golang standard libraries

## to run server
```
  go run main.go
```

## to upload image 
```
  curl localhost:9091/images/{id}/{filename} -d {file_path}
```
### example
```
  curl localhost:9091/images/{id}/newfilename.png -d @test.png
```


## to retrieve image 
```
  curl localhost:9091/images/{id}/{filename}
```
### example
```
  curl localhost:9091/images/{id}/newfilename.png
```
