package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(directionCmd)
}

var directionCmd = &cobra.Command{
	Use:   "directions",
	Short: "Go from point A to B",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🌝 Directions")
	},
}