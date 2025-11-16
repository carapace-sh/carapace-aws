package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_getAccountSummaryCmd = &cobra.Command{
	Use:   "get-account-summary",
	Short: "Retrieves information about IAM entity usage and IAM quotas in the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_getAccountSummaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_getAccountSummaryCmd).Standalone()

	})
	iamCmd.AddCommand(iam_getAccountSummaryCmd)
}
