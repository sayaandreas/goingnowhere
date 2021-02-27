package main

import "github.com/sayaandreas/goingnowhere/cmd"

func main() {
	cmd.Execute()
}

// func multipleFiles() {
// 	fmt.Println("Download Started")
// 	urls := []string{"https://upload.wikimedia.org/wikipedia/commons/thumb/0/0a/Ambarawa_Bypass_Road_from_Eling_Bening%2C_2017-03-15.jpg/3240px-Ambarawa_Bypass_Road_from_Eling_Bening%2C_2017-03-15.jpg", "https://upload.wikimedia.org/wikipedia/commons/d/d6/Wp-w4-big.jpg"}
// 	bytes, err := download.DownloadMultipleFilesIntoBytes(urls)
// 	if err != nil {
// 		panic(err)
// 	}
// 	for i := 0; i < len(urls); i++ {
// 		output, err := os.Create("files" + strconv.Itoa(i) + ".jpg")
// 		if err != nil {
// 			panic(err)
// 		}
// 		size, err := output.Write(bytes[i])
// 		fmt.Printf("wrote %d bytes\n", size)
// 	}
// }

// func withProgress() {
// 	fmt.Println("Download Started")

// 	fileURL := "https://upload.wikimedia.org/wikipedia/commons/d/d6/Wp-w4-big.jpg"
// 	err := download.DownloadFileWithProgress("avatar.jpg", fileURL)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Download Finished")
// }
