# sumdiff
Compare sdk/tools for directory and files.

# SDK

## Install
```bash
go get -u github.com/foolin/sumdiff
```

## API
```go

//Compare 
ok, result, err := Cmp(path1, path2)

//Equal
ok, err := Equal(v.path1, v.path2)

```

# CLI Tool

```
sumdiff  --help
A useful comparison tool for differences and hash

Usage:
  sumdiff [command]

Available Commands:
  eq          Compare whether two files or directory are equal
  cmp         Compare the two files or directories are different
  hash        Calculate hash algorithm [md5|sha1|sha256|sha512] hex string
  help        Help about any command
  md5         Calculate md5 hex string
  sha1        Calculate sha1 hex string
  sha256      Calculate sha256 hex string
  completion  Generate the autocompletion script for the specified shell
  
Flags:
  -h, --help      help for sumdiff
  -v, --verbose   Verbose output info

Use "sumdiff [command] --help" for more information about a command.

```

# Usage

### `sumdiff eq` - Compare equal
* Compare the two files or directories are equal
  `sumdiff eq <path1> <path2>`

  Example1:
    ```shell
    sumdiff eq test_data/a.txt test_data/b.txt 
    ```
  Output1:
    ```text                                                                               
    true
    ```
    
  Example2:
    ```shell
    sumdiff eq test_data/c.txt test_data/d.txt 
    ```
  Output2:
    ```text                                                                                                                                                                    
    false
    ```



### `sumdiff cmp` - Compare different true or false
* Compare the two file or directory differences
  `sumdiff diff <path1> <path2>`

Example:
  ```shell
  sumdiff cmp test_data/data1 test_data/data3
  ```
Output:
  ```text                                                                               
  +-------------------------------------------+                                                       
  | Path   | OK    | Msg                      |
  +-------------------------------------------+
  |        | true  |                          |
  | /a.txt | true  |                          |
  | /b.txt | true  |                          |
  | /c.txt | false | path1 not exist [/c.txt] |
  +-------------------------------------------+
  false
  ```


### `sumdiff md5` - Get MD5 hash value
  Calculate the MD5 hexadecimal value of the files or directories
  `sumdiff md5 <path1> [<path2>] [<path3>] ...`
  
  Example:
  ```shell
  sumdiff md5 test_data/a.txt
  ```
  Output:
  ```text                                                                               
  +-----------------------------------------------------------+                                       
  | Hash                             | Size | Path            |
  +-----------------------------------------------------------+
  | 9d15fa011b54dbd079d1d20e36e4a358 | 212  | test_data/a.txt |
  +-----------------------------------------------------------+
  ```

### `sumdiff sha1` - Get SHA1 hash value
  Calculate the MD5 hexadecimal value of the files or directories
  `sumdiff sha1 <path1> [<path2>] [<path3>] ...`
  
  Example:
  ```shell
  sumdiff sha1 test_data/data1
  ```
  Output:
  ```text                                                                               
  +---------------------------------------------------------------+                                   
  | Hash                                     | Size | Path        |
  +---------------------------------------------------------------+
  | 18b012c3e30bf822589ac96fd4d87e6e8d89754e | 4096 | data1/a.txt |
  | 9c678f3b44c0918f5695e5a0a8232ab1b017a4a4 | 4096 | data1/b.txt |
  +---------------------------------------------------------------+
  ```

### `sumdiff sha256` - Get SHA256 hash value
  `sumdiff sha256 <path1> [<path2>] [<path3>] ...`
  
  Example:
  ```shell
  sumdiff sha256 test_data/a.txt test_data/c.txt
  ```
  Output:
  ```text                                                                               
  +-------------------------------------------------------------------------------------------+       
  | Hash                                                             | Size | Path            |
  +-------------------------------------------------------------------------------------------+
  | bb3cf386d0c975847024b0d78c9f92c0657c894f16c1db1248233cfe2b05b65f | 212  | test_data/a.txt |
  | 732d2d9a2c39480fa38bd060f07b3e55eb3b47888e8019a8d4cfd8600d5c104f | 222  | test_data/c.txt |
  +-------------------------------------------------------------------------------------------+
  ```

### `sumdiff hash` - Get Other hash value 
  Calculate the  hash(md5|sha1|sha256|sha512) value of a file or directory
  `sumdiff hash <md5|sha1|sha256|sha512> <path1> [<path2>] [<path3>] ...`
  
  Example:
  ```shell
  sumdiff hash sha512 test_data/a.txt
  ```
  Output:
  ```text                                                                               
  +-----------------------------------------------------------------------------------------------------------------------------------------------------------+
  | Hash                                                                                                                             | Size | Path            |
  +-----------------------------------------------------------------------------------------------------------------------------------------------------------+
  | 9683aedddcc5b16548f3510580d91306fad405070fd516299e3f5609bff5fd950a1a6e39e8bce5000d4f3c264428855eb2ae0f235f55d89bd9ec2c9f02c86c4b | 212  | test_data/a.txt |
  +-----------------------------------------------------------------------------------------------------------------------------------------------------------+
  ```

## Help
`sumdiff -h`


## Install

linux:
```shell

tar -xvf sumdiff_Linux_x86_64.tar.gz && sudo mv sumdiff /usr/local/bin

```