package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_generateCredentialReportCmd = &cobra.Command{
	Use:   "generate-credential-report",
	Short: "Generates a credential report for the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_generateCredentialReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_generateCredentialReportCmd).Standalone()

	})
	iamCmd.AddCommand(iam_generateCredentialReportCmd)
}
