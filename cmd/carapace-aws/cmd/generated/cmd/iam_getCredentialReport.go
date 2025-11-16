package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_getCredentialReportCmd = &cobra.Command{
	Use:   "get-credential-report",
	Short: "Retrieves a credential report for the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_getCredentialReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_getCredentialReportCmd).Standalone()

	})
	iamCmd.AddCommand(iam_getCredentialReportCmd)
}
