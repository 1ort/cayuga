# cayuga
Cayuga is regexp-based TON address generator
![изображение](https://user-images.githubusercontent.com/83316072/209030972-9d62f403-6ac1-425f-9ae9-00363d86c0a8.png)


## Installation
via git:

```
git clone https://github.com/1ort/cayuga.git
cd cayuga
go build .
```


## Usage

 ```cayuga -r hi$ -i -w 10 ```

```
-i    Ignore letter case
-r string
    Regexp pattern (default "^")
-w int
    Number of workers (default 1)
```
Use ctrl+c to stop the program, otherwise it will iterate indefinitely
