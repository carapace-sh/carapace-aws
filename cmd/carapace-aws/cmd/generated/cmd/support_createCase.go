package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var support_createCaseCmd = &cobra.Command{
	Use:   "create-case",
	Short: "Creates a case in the Amazon Web Services Support Center.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(support_createCaseCmd).Standalone()

	support_createCaseCmd.Flags().String("attachment-set-id", "", "The ID of a set of one or more attachments for the case.")
	support_createCaseCmd.Flags().String("category-code", "", "The category of problem for the support case.")
	support_createCaseCmd.Flags().String("cc-email-addresses", "", "A list of email addresses that Amazon Web Services Support copies on case correspondence.")
	support_createCaseCmd.Flags().String("communication-body", "", "The communication body text that describes the issue.")
	support_createCaseCmd.Flags().String("issue-type", "", "The type of issue for the case.")
	support_createCaseCmd.Flags().String("language", "", "The language in which Amazon Web Services Support handles the case.")
	support_createCaseCmd.Flags().String("service-code", "", "The code for the Amazon Web Services service.")
	support_createCaseCmd.Flags().String("severity-code", "", "A value that indicates the urgency of the case.")
	support_createCaseCmd.Flags().String("subject", "", "The title of the support case.")
	support_createCaseCmd.MarkFlagRequired("communication-body")
	support_createCaseCmd.MarkFlagRequired("subject")
	supportCmd.AddCommand(support_createCaseCmd)
}
