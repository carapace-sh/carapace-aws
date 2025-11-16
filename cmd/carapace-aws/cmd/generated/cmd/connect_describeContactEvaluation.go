package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeContactEvaluationCmd = &cobra.Command{
	Use:   "describe-contact-evaluation",
	Short: "Describes a contact evaluation in the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeContactEvaluationCmd).Standalone()

	connect_describeContactEvaluationCmd.Flags().String("evaluation-id", "", "A unique identifier for the contact evaluation.")
	connect_describeContactEvaluationCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_describeContactEvaluationCmd.MarkFlagRequired("evaluation-id")
	connect_describeContactEvaluationCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_describeContactEvaluationCmd)
}
