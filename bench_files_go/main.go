package main

import (
    "log"
    "net/http"
    "os"
    "io"
    "strconv"
)

func handler(writer http.ResponseWriter, r *http.Request) {
    //Check if file exists and open
    Filename := "/home/dyens/dev/bench_files/my_file"
    Openfile, err := os.Open(Filename)
    defer Openfile.Close() //Close after function return
    if err != nil {
    	//File not found, send 404
    	http.Error(writer, "File not found.", 404)
    	return
    }
    
    //File is found, create and send the correct headers
    
    //Get the Content-Type of the file
    //Create a buffer to store the header of the file in
    FileHeader := make([]byte, 512)
    //Copy the headers into the FileHeader buffer
    Openfile.Read(FileHeader)
    //Get content type of file
    FileContentType := http.DetectContentType(FileHeader)
    
    //Get the file size
    FileStat, _ := Openfile.Stat()                     //Get info from file
    FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string
    
    //Send the headers
    writer.Header().Set("Content-Disposition", "attachment; filename="+Filename)
    writer.Header().Set("Content-Type", FileContentType)
    writer.Header().Set("Content-Length", FileSize)
    
    //Send the file
    //We read 512 bytes from the file already, so we reset the offset back to 0
    Openfile.Seek(0, 0)
    io.Copy(writer, Openfile) //'Copy' the file to the client
    return
}



func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}