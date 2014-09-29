package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "os"
    "bufio"
    "strconv"
)

type Result struct {
    min_w string
    max_w string
    min_h string
    max_h string
    value string
}

const INPUT_FILE = "src.txt"
const OUTPUT_FILE = "dst.txt"
const START_POSITION = 2

func main() {
    var rfp, wfp *os.File
    var err error

    // ファイルオープン
    rfp, err = os.Open(INPUT_FILE)
    if err != nil {
        fmt.Println("file open error")
        return
    }
    defer rfp.Close()

    // ファイル読込
    reader := csv.NewReader(rfp)
    reader.Comma = '\t'
    var read_data [][]string
    for {
       record, err := reader.Read()
       if err == io.EOF {
           break
       } else if err != nil {
           fmt.Println("file read error")
           return
       }
       read_data = append(read_data, record)
    }

    // 構造体(Result)設定
    var result_list []*Result
    for i := START_POSITION; i < len(read_data); i++ {
        for j := START_POSITION; j < len(read_data[i]); j++ {
            result := new(Result)
            min_w := read_data[0][j - 1]
            if j != START_POSITION {
                var min_w_i int
                min_w_i, _ = strconv.Atoi(min_w)
                min_w_i = min_w_i + 1
                min_w = strconv.Itoa(min_w_i)
            }
            result.min_w = min_w
            result.max_w = read_data[0][j]
            min_h := read_data[i - 1][0]
            if i != START_POSITION {
                var min_h_i int
                min_h_i, _ = strconv.Atoi(min_h)
                min_h_i = min_h_i + 1
                min_h = strconv.Itoa(min_h_i)
            }
            result.min_h = min_h
            result.max_h = read_data[i][0]
            result.value = read_data[i][j]
            result_list = append(result_list, result)
        }
    }

    // ファイル出力
    wfp, err = os.OpenFile(OUTPUT_FILE, os.O_WRONLY | os.O_CREATE, 0600)
    if err != nil {
        fmt.Println("file open error")
        return
    }
    defer wfp.Close()
    err = wfp.Truncate(0)
    if err != nil {
        fmt.Println("file truncate error")
        return
    }
    writer := bufio.NewWriter(wfp)
    for i := 0; i < len(result_list); i++ {
        writer.WriteString(result_list[i].min_w + "\t" +
                      result_list[i].max_w + "\t" +
                      result_list[i].min_h + "\t" +
                      result_list[i].max_h + "\t" +
                      result_list[i].value + "\n")
    }
    writer.Flush()

    fmt.Println("output to " + OUTPUT_FILE)
}
