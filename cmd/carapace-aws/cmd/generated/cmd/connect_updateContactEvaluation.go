package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateContactEvaluationCmd = &cobra.Command{
	Use:   "update-contact-evaluation",
	Short: "Updates details about a contact evaluation in the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateContactEvaluationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_updateContactEvaluationCmd).Standalone()

		connect_updateContactEvaluationCmd.Flags().String("answers", "", "A map of question identifiers to answer value.")
		connect_updateContactEvaluationCmd.Flags().String("evaluation-id", "", "A unique identifier for the contact evaluation.")
		connect_updateContactEvaluationCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_updateContactEvaluationCmd.Flags().String("notes", "", "A map of question identifiers to note value.")
		connect_updateContactEvaluationCmd.Flags().String("updated-by", "", "The ID of the user who updated the contact evaluation.")
		connect_updateContactEvaluationCmd.MarkFlagRequired("evaluation-id")
		connect_updateContactEvaluationCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_updateContactEvaluationCmd)
}
