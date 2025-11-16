package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_batchGetAccountStatusCmd = &cobra.Command{
	Use:   "batch-get-account-status",
	Short: "Retrieves the Amazon Inspector status of multiple Amazon Web Services accounts within your environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_batchGetAccountStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_batchGetAccountStatusCmd).Standalone()

		inspector2_batchGetAccountStatusCmd.Flags().String("account-ids", "", "The 12-digit Amazon Web Services account IDs of the accounts to retrieve Amazon Inspector status for.")
	})
	inspector2Cmd.AddCommand(inspector2_batchGetAccountStatusCmd)
}
