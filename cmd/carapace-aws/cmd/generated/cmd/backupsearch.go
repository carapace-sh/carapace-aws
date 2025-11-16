package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupsearchCmd = &cobra.Command{
	Use:   "backupsearch",
	Short: "Backup Search\n\nBackup Search is the recovery point and item level search for Backup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupsearchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backupsearchCmd).Standalone()

	})
	rootCmd.AddCommand(backupsearchCmd)
}
