package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_getLegalHoldCmd = &cobra.Command{
	Use:   "get-legal-hold",
	Short: "This action returns details for a specified legal hold.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_getLegalHoldCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_getLegalHoldCmd).Standalone()

		backup_getLegalHoldCmd.Flags().String("legal-hold-id", "", "The ID of the legal hold.")
		backup_getLegalHoldCmd.MarkFlagRequired("legal-hold-id")
	})
	backupCmd.AddCommand(backup_getLegalHoldCmd)
}
