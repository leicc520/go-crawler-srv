package proxy

import (
	"compress/gzip"
	"errors"
	"github.com/leicc520/go-orm/log"
	"golang.org/x/text/encoding/charmap"
	"io"
	"io/ioutil"
)

//对反馈回来的内容gzip解压
func GZIPDecode(reader io.Reader) (bytes []byte, err error) {
	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return
	}
	defer func() {
		if err = gzipReader.Close(); err != nil {
			log.Write(log.ERROR, "gzip reader close fail", err)
		}
	}()
	bytes, err = ioutil.ReadAll(gzipReader)
	return
}

//根据编码类别做信息解码处理
func Decode(encoding string, bytes []byte) (s string, err error) {
	var char *charmap.Charmap
	if encoding == "iso-8859-1" {
		char = charmap.ISO8859_1
	} else {
		err = errors.New("要解码的编码不支持:"+encoding)
		return
	}

	bytes, err = char.NewDecoder().Bytes(bytes)
	if err != nil {
		return
	}
	s = string(bytes)
	return
}