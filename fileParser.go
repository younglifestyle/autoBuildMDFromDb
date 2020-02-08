package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
	"io/ioutil"
	"os"
	"regexp"
)

const (
	//分离表
	tablePattern = `CREATE TABLE([\s\S]*?);`

	// 获取表名称
	tableNamePattern = "CREATE TABLE `(.*)`"
	// 获取字段名称
	namePattern = "`(.*?)` ([\\S]{1,}).* [COMMENT]{0,} ('(.*?)'){0,}"
	//获取表详情
	tableComment = `COMMENT=\'(.*?)\'`
)

var (
	regTableContent     = regexp.MustCompile(tablePattern)
	regTableNamePattern = regexp.MustCompile(tableNamePattern)
	regNamePattern      = regexp.MustCompile(namePattern)
	regTableComment     = regexp.MustCompile(tableComment)

	fileName = "machine.sql"
)

func main() {
	fileName1 := flag.String("f", "machine.sql", "SQL文件")
	flag.Parse()

	fileName = *fileName1
	tableContentList := separatTable()

	tableResMap := make(map[string][][]string)
	tableNameList := make([]string, 0)
	requiredKeynameAndComment(tableContentList,
		tableResMap, &tableNameList)

	fi, err := os.OpenFile(fileName+".md",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Printf("open file error :%v \n", err)
		return
	}

	write2Md(fi, tableResMap, tableNameList)
	defer fi.Close()
}

func write2Md(writer io.Writer,
	tableResMap map[string][][]string,
	tableNameList []string) error {

	writer.Write([]byte("# 数据库文档\n\n"))
	for _, tableName := range tableNameList {
		tableKeyMap := tableResMap[tableName]
		sprintf := fmt.Sprintf("%s表", tableName)
		if len(tableKeyMap[0]) != 0 {
			sprintf = fmt.Sprintf("%s(%s)",
				sprintf, tableKeyMap[0][0])
		}

		writer.Write([]byte("## " + sprintf + "\n\n"))

		table := tablewriter.NewWriter(writer)
		table.SetHeader([]string{"字段名称", "字段类型", "字段含义"})
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")
		table.AppendBulk(tableKeyMap[1:]) // Add Bulk Data
		table.Render()
	}

	return nil
}

func requiredKeynameAndComment(tableContentList [][]byte,
	tableRes map[string][][]string, tableNameList *[]string) {
	for _, tableContent := range tableContentList {
		contentTmp := bytes.Split(tableContent, []byte("\n"))
		tableName := regTableNamePattern.FindSubmatch(contentTmp[0])
		tableComment := regTableComment.FindSubmatch(contentTmp[len(contentTmp)-1])

		var keyArr [][]string
		if len(tableComment) > 1 {
			keyArr = [][]string{{string(tableComment[1])}} // 第一个位置留给table comment
		} else {
			keyArr = [][]string{{}} // 第一个位置留给table comment
		}

		//  key  key_type key_comment(include)
		for _, contentLine := range contentTmp[1 : len(contentTmp)-1] {
			keyName := regNamePattern.FindSubmatch(contentLine)

			if len(keyName) > 2 {
				keyTmp := []string{string(keyName[1]), string(keyName[2])}
				if len(keyName) > 3 || string(keyName[3]) != "" {
					keyTmp = append(keyTmp, string(keyName[4]))
				}
				keyArr = append(keyArr, keyTmp)
			}
		}

		tableRes[string(tableName[1])] = keyArr // 新建表
		*tableNameList = append(*tableNameList, string(tableName[1]))
	}
}

func separatTable() [][]byte {
	readBy, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return regTableContent.FindAll(readBy, -1)
}
