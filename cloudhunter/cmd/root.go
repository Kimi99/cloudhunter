package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cloudhunter",
	Short: "CloudHunter - AWS post-compromise enumeration tool",
	Long: `CloudHunter is a CLI tool for mapping AWS environments using 
stolen or assumed credentials in post-compromise or red team scenarios.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}