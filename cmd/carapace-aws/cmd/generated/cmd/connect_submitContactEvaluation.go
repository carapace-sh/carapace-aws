package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_submitContactEvaluationCmd = &cobra.Command{
	Use:   "submit-contact-evaluation",
	Short: "Submits a contact evaluation in the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_submitContactEvaluationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_submitContactEvaluationCmd).Standalone()

		connect_submitContactEvaluationCmd.Flags().String("answers", "", "A map of question identifiers to answer value.")
		connect_submitContactEvaluationCmd.Flags().String("evaluation-id", "", "A unique identifier for the contact evaluation.")
		connect_submitContactEvaluationCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_submitContactEvaluationCmd.Flags().String("notes", "", "A map of question identifiers to note value.")
		connect_submitContactEvaluationCmd.Flags().String("submitted-by", "", "The ID of the user who submitted the contact evaluation.")
		connect_submitContactEvaluationCmd.MarkFlagRequired("evaluation-id")
		connect_submitContactEvaluationCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_submitContactEvaluationCmd)
}
