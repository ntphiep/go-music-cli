package cmd

import (
	"fmt"
	"time"

	"github.com/ntphiep/go-music-cli/counter"
	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Count the number of lines in a file",
	Long: `A long description
	that spans multiple lines
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to count the number of lines in a file.`,

	Run: func(cmd *cobra.Command, args []string) {

		// Check args
		if len(args) == 1 {
			start := time.Now()

			// Call the read and calculate function
			outs, err := counters.CountOthers(args[0])
			if err != nil {
				fmt.Println(err)

			} else {
				arr := []string{"vowel", "consonant", "space", "digit", "pmark", "letter"}
				for order, f := range arr {
					fstatus, _ := cmd.Flags().GetBool(f)

					if fstatus {
						fmt.Printf("%s count: %d\n", f, outs[order])
					}
				}
			}

			// Call the word count function
			resutls, err := counters.CountWord(args[0])
			if err != nil {
				fmt.Println(err)
			} else {
				fstatus, _ := cmd.Flags().GetBool("word")
				if fstatus {
					fmt.Printf("word count: %d\n", resutls)
				}
			}

			// Check for time flag
			fstatus, _ := cmd.Flags().GetBool("time")
			if fstatus {
				fmt.Printf("Time taken: %v\n", time.Since(start))
			}
		} else {
			fmt.Println("Please provide a file path")
		}
	},
}

func init() {

}
