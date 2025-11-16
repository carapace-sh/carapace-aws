package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeEvaluationFormCmd = &cobra.Command{
	Use:   "describe-evaluation-form",
	Short: "Describes an evaluation form in the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeEvaluationFormCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_describeEvaluationFormCmd).Standalone()

		connect_describeEvaluationFormCmd.Flags().String("evaluation-form-id", "", "A unique identifier for the contact evaluation.")
		connect_describeEvaluationFormCmd.Flags().String("evaluation-form-version", "", "A version of the evaluation form.")
		connect_describeEvaluationFormCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_describeEvaluationFormCmd.MarkFlagRequired("evaluation-form-id")
		connect_describeEvaluationFormCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_describeEvaluationFormCmd)
}
