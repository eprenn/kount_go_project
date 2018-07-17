package main

import ("fmt"
        "os"
        "github.com/fatih/color")
// function to print all files in a directory with some error handling if there is an issue opening or reading a directory.
func print_dir(fpath string, supress_stdout bool) (err error, stdout string) {
    
    var standard_out = ""
    directory, err := os.Open(fpath)
    defer directory.Close()
    if err != nil {
        fmt.Println(err)
        return err, standard_out
    }
    
    file_info, err := directory.Readdir(0)
    if err != nil {
        fmt.Println(err)
        return err, standard_out
    }
    
    my_color := color.New(color.FgCyan)
    for _, f := range file_info {
        if f.Mode().IsRegular() {
            standard_out += f.Name() + "\n"
            if !supress_stdout {
                my_color.Println(f.Name())
            }
        }
    }
    return err, standard_out
}

func main() {
	print_dir("../my_func", false)
}

 
