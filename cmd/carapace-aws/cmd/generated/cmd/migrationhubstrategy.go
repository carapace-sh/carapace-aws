package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubstrategyCmd = &cobra.Command{
	Use:   "migrationhubstrategy",
	Short: "Migration Hub Strategy Recommendations\n\nThis API reference provides descriptions, syntax, and other details about each of the actions and data types for Migration Hub Strategy Recommendations (Strategy Recommendations).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubstrategyCmd).Standalone()

	rootCmd.AddCommand(migrationhubstrategyCmd)
}
