package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listLegalHoldsCmd = &cobra.Command{
	Use:   "list-legal-holds",
	Short: "This action returns metadata about active and previous legal holds.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listLegalHoldsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_listLegalHoldsCmd).Standalone()

		backup_listLegalHoldsCmd.Flags().String("max-results", "", "The maximum number of resource list items to be returned.")
		backup_listLegalHoldsCmd.Flags().String("next-token", "", "The next item following a partial list of returned resources.")
	})
	backupCmd.AddCommand(backup_listLegalHoldsCmd)
}
