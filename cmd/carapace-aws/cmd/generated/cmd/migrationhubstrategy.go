package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubstrategyCmd = &cobra.Command{
	Use:   "migrationhubstrategy",
	Short: "Migration Hub Strategy Recommendations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubstrategyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhubstrategyCmd).Standalone()

	})
	rootCmd.AddCommand(migrationhubstrategyCmd)
}
