package eg4

import "testing"

func TestAll(t *testing.T) {

	var downloader Downloader = NewHTTPDownloader()
	downloader.Download("http://example.com/abc.zip")

	downloader = NewFTPDownloader()
	downloader.Download("ftp://example.com/abc.zip")
}
