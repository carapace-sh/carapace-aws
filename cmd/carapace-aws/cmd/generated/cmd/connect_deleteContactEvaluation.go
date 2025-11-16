package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteContactEvaluationCmd = &cobra.Command{
	Use:   "delete-contact-evaluation",
	Short: "Deletes a contact evaluation in the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteContactEvaluationCmd).Standalone()

	connect_deleteContactEvaluationCmd.Flags().String("evaluation-id", "", "A unique identifier for the contact evaluation.")
	connect_deleteContactEvaluationCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_deleteContactEvaluationCmd.MarkFlagRequired("evaluation-id")
	connect_deleteContactEvaluationCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_deleteContactEvaluationCmd)
}
