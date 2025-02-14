package main

import (
	"encoding/csv"
	"os"
)

func main()  {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"m", "fahmi", "syaifudin"})
	_ = writer.Write([]string{"mohamad", "fahmi", "syaifudin"})
	
	writer.Flush()
}