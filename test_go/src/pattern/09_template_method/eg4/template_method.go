package eg4

import "fmt"

//下载器接口
type Downloader interface {
	Download(uri string)
}

//下载工具接口
type Implement interface {
	download()
	save()
}

/*===================================================================================================================*/

//模板
type Template struct {
	Implement
	uri string
}

func NewTemplate(impl Implement) *Template {
	return &Template{
		Implement: impl,
		uri:       "",
	}
}

//实现下载器的接口（这是一个模板方法）
func (this *Template) Download(uri string) {
	this.uri = uri
	fmt.Println("prepare downloading")
	this.Implement.download()
	this.Implement.save()
	fmt.Println("finish downloading")
}

//实现工具的save函数
func (this *Template) save() {
	fmt.Println("finish downloading")
}

/*====================================================================================================================*/

//1.http下载器
type HTTPDownloader struct {
	*Template
}

func NewHTTPDownloader() Downloader {
	this := new(HTTPDownloader)
	this.Template = NewTemplate(this)
	return this
}

func (this *HTTPDownloader) download() {
	fmt.Printf("download %s via http\n", this.uri)
}

func (this *HTTPDownloader) save() {
	fmt.Printf("http save\n")
}

//2.FTP下载器
type FTPDownloader struct {
	*Template
}

func NewFTPDownloader() Downloader {
	this := new(FTPDownloader)
	this.Template = NewTemplate(this)
	return this
}

func (this *FTPDownloader) download() {
	fmt.Printf("download %s via ftp\n", this.uri)
}
