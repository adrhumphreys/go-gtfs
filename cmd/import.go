package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"gtfs-cli/internal"
	"gtfs-cli/internal/dbx"
	"os"
)

func init() {
	rootCmd.AddCommand(importCmd)
}

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import GTFS data from a source",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a folder argument to import from")
		}

		_, err := os.Stat(args[0])
		if os.IsNotExist(err) {
			return fmt.Errorf("the folder specified doesn't exist: %v\n", args[0])
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚇 Importing")
		dbx.InitDB()
		internal.Import(args[0])
	},
}