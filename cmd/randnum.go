/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

// randnumCmd represents the randnum command
var randnumCmd = &cobra.Command{
	Use:   "randnum",
	Short: "random 10 numbers between 0-100",
	Long:  `config able via flag`,

	Run: func(cmd *cobra.Command, args []string) {

		echo, _ := cmd.Flags().GetBool("echo")
		count, _ := cmd.Flags().GetInt("count")
		minBound, _ := cmd.Flags().GetInt("start")
		maxBound, _ := cmd.Flags().GetInt("end")

		if echo {
			fmt.Println(count)
		}
		for i := 0; i < int(count); i++ {
			randomed := rand.Intn(int(maxBound-minBound+1)) + int(minBound)
			fmt.Print(randomed, " ")
		}
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(randnumCmd)

	randnumCmd.Flags().BoolP("echo", "n", false, "also output amount of number")
	randnumCmd.Flags().IntP("count", "i", 10, "amount of number to generate")
	randnumCmd.Flags().IntP("start", "s", 0, "number start at")
	randnumCmd.Flags().IntP("end", "e", 100, "number end at (included)")

}
