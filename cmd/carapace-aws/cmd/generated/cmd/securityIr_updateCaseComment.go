package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityIr_updateCaseCommentCmd = &cobra.Command{
	Use:   "update-case-comment",
	Short: "Updates an existing case comment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityIr_updateCaseCommentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityIr_updateCaseCommentCmd).Standalone()

		securityIr_updateCaseCommentCmd.Flags().String("body", "", "Required element for UpdateCaseComment to identify the content for the comment to be updated.")
		securityIr_updateCaseCommentCmd.Flags().String("case-id", "", "Required element for UpdateCaseComment to identify the case ID containing the comment to be updated.")
		securityIr_updateCaseCommentCmd.Flags().String("comment-id", "", "Required element for UpdateCaseComment to identify the case ID to be updated.")
		securityIr_updateCaseCommentCmd.MarkFlagRequired("body")
		securityIr_updateCaseCommentCmd.MarkFlagRequired("case-id")
		securityIr_updateCaseCommentCmd.MarkFlagRequired("comment-id")
	})
	securityIrCmd.AddCommand(securityIr_updateCaseCommentCmd)
}
