package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gtfs-cli/internal"
)

func init() {
	rootCmd.AddCommand(routeCmd)
}

var routeCmd = &cobra.Command{
	Use:   "routes",
	Short: "List all routes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚃 Routes")
		routes := internal.GetRoutes()

		for _, route := range routes {
			if route.Type == 3 {
				continue
			}

			fmt.Printf("%v - %v - %v - %v\n",
				route.TypeEmoji(),
				route.ID,
				route.ShortName,
				route.LongName)
		}
	},
}
