# Mockserver

A Go tool that mocks HTTP endpoints.

You can specify response formats for an endpoint and use `mockerserver` to serve those responses. This tool is meant to be used for testing purposes.

## Install

Required
- Go 1.8 and above

``` bash
$ go get github.com/fanglinli/mockserver
$ go build
```

## Usage

```
Usage of ./mockserver:
  -folder string
    	Folder location of response specification file. (default ".")
  -listen string
    	Port to listen on. (default ":8080")
```

Each file in the target folder will create an endpoint with the filename as the endpoint name. Each endpoint will serve a response with a response body that corresponds to the file content's format and a `200` status code. Standard `go` string format specifiers can be used to return random values that match the specified types.

### Supported types:

|Type|Format String|Specifying Size|
|---|---|---|
|String|`%s`|Use an integer in the middle of the format string to specify string length. Eg. `%8s` to specify a string of length 8.|
|Number|`%d`|Use an integer in the middle of the format string to specify number of digits. Eg. `%3s` to specify a number with 3 digits.|

Note: Default size is five.

## Example

``` bash
$ ./mockserver -folder example
$ curl localhost:8080/employees
```

Returns a response like the following:

``` json
{
    "locationID": "1947779410",
    "employees": [
        {
            "name": "bwbmv ixrzj",
            "id": "4263669287",
            "email": "zlqxtrqgxg@lmdjy.com",
            "phone": "+1 7193831650"
        },
        {
            "name": "nwutx jdoih",
            "id": "7808658932",
            "email": "fqrvwchumk@yfwsm.com",
            "phone": "+1 8050264449"
        },
        {
            "name": "eiawj tbekz",
            "id": " 125624717",
            "email": "ueuahxmcfh@tbutu.com",
            "phone": "+1 4598140041"
        }
    ]
}
```

## License

Copyright (c) 2017 Fang Lin Li

Please consider promoting this project if you find it useful.

Permission is hereby granted, free of charge, to any person
obtaining a copy of this software and associated documentation
files (the "Software"), to deal in the Software without restriction,
including without limitation the rights to use, copy, modify, merge,
publish, distribute, sublicense, and/or sell copies of the Software,
and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included
in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT
OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
