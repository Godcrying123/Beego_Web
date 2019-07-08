package utils

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

func init() {
	currentdate := time.Now().Format("2006-1-2")
	logfilepath := "logs/" + currentdate + ".log"
	var loglevel int
	loglevel, _ = strconv.Atoi(beego.AppConfig.String("loglevel"))
	_, err := os.OpenFile(logfilepath, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		beego.Error(err)
	}
	//var logfileconfig string = """filename":"logfilepath"
	var logfileconfig string = "{\"filename\":\"" + logfilepath + "\"}"
	beego.SetLogger("file", logfileconfig)
	//beego.BeeLogger.DelLogger("console")
	beego.SetLevel(loglevel)
}

func AllLogsGet() {
	var logpath string = "logs/"
	_, err := os.Stat(logpath)
	//beego.Info(err)
	if os.IsExist(err) {
		beego.Error("logs folder not existed!")
		return
	} else {
		beego.Info("logs folder existed!")
	}
	alllogfiles, err1 := ioutil.ReadDir(logpath)
	if err1 != nil {
		beego.Error(err1)
	} else {
		err1 = RankOldLogsByTime(alllogfiles, 2)
		if err1 != nil {
			beego.Error(err1)
		}
	}

}

func RankOldLogsByTime(fileinfos []os.FileInfo, archievenum int) (err error) {
	var filenames []string
	//var deletelog string
	for _, singlelogfile := range fileinfos {
		//filenames = filenames[:0]
		filenames = append(filenames, singlelogfile.Name())
		//beego.Info(len(filenames))
	}
	if len(filenames) > 10 {
		filenames = filenames[:archievenum]
		zipname := strings.Join(filenames[1:], "")
		//fmt.Println(zipname)
		err = ZipFiles(zipname, filenames)
		return err
	}
	return nil
}

func ZipFiles(filename string, files []string) error {
	var zipfile string = "logs/" + filename + ".zip"
	var newzipfile string = "archievelogs/" + filename + ".zip"
	newZipFile, err := os.Create(zipfile)
	if err != nil {
		beego.Error(err)
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {
		if err = AddFileToZip(zipWriter, file); err != nil {
			beego.Error(err)
		}
	}
	err = os.Rename(zipfile, newzipfile)
	if err != nil {
		beego.Error(err)
	}
	return nil
}

func AddFileToZip(zipWriter *zip.Writer, filename string) error {

	fileToZip, err := os.Open("logs/" + filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// Using FileInfoHeader() above only uses the basename of the file. If we want
	// to preserve the folder structure we can overwrite this with the full path.
	header.Name = filename

	// Change to deflate to gain better compression
	// see http://golang.org/pkg/archive/zip/#pkg-constants
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
