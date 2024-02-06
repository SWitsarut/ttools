package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var issortCmd = &cobra.Command{
	Use:   "issort",
	Short: "Check if a list of numbers provided via stdin is sorted",
	Long:  `This command checks if a list of numbers provided via stdin is sorted.`,
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)

		first := true
		var prev float64
		for scanner.Scan() {
			line := scanner.Text()
			numStrs := strings.Fields(line)

			for _, numStr := range numStrs {
				num, err := strconv.ParseFloat(numStr, 64)
				if err != nil {
					panic("cannot convert input into number")
				}
				if !first {
					if prev > num {
						fmt.Println(false)
						os.Exit(0)
					}
				} else {
					first = false
				}
				prev = num
			}
		}

		if !first {
			fmt.Println(true)
			os.Exit(0)
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(issortCmd)
}
