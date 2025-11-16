package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupsearchCmd = &cobra.Command{
	Use:   "backupsearch",
	Short: "Backup Search",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupsearchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backupsearchCmd).Standalone()

	})
	rootCmd.AddCommand(backupsearchCmd)
}
