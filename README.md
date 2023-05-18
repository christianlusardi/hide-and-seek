![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
# Hide & Seek

An academic simple tool for testing stenography in Go



## Run Locally

Clone the project

```bash
  git clone https://github.com/christianlusardi/hide-and-seek
```

Go to the project directory

```bash
  cd hide-and-seek
```

Build the project

```bash
  go build
```

Or run wihtout native compilation

```bash
  go run main.go
```




## Usage/Examples

Encode (hide a text from file into an image)

```bash
./hideandseek -e -mi my-txt-location -pi my-input-image-location.png -po my-output-img-location.png
```

Decode (read a text from image)


```bash
./hideandseek -d -pi my-image-with-data-location.png
```


Help

```bash
./hideandseek -help
```
## Authors

- [@christianlusardi](https://www.github.com/christianlusardi)

