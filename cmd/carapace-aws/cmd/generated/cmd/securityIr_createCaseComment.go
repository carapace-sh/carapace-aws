package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityIr_createCaseCommentCmd = &cobra.Command{
	Use:   "create-case-comment",
	Short: "Adds a comment to an existing case.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityIr_createCaseCommentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityIr_createCaseCommentCmd).Standalone()

		securityIr_createCaseCommentCmd.Flags().String("body", "", "Required element used in combination with CreateCaseComment to add content for the new comment.")
		securityIr_createCaseCommentCmd.Flags().String("case-id", "", "Required element used in combination with CreateCaseComment to specify a case ID.")
		securityIr_createCaseCommentCmd.Flags().String("client-token", "", "The `clientToken` field is an idempotency key used to ensure that repeated attempts for a single action will be ignored by the server during retries.")
		securityIr_createCaseCommentCmd.MarkFlagRequired("body")
		securityIr_createCaseCommentCmd.MarkFlagRequired("case-id")
	})
	securityIrCmd.AddCommand(securityIr_createCaseCommentCmd)
}
