package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_updateAnswerCmd = &cobra.Command{
	Use:   "update-answer",
	Short: "Update the answer to a specific question in a workload review.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_updateAnswerCmd).Standalone()

	wellarchitected_updateAnswerCmd.Flags().String("choice-updates", "", "A list of choices to update on a question in your workload.")
	wellarchitected_updateAnswerCmd.Flags().String("is-applicable", "", "")
	wellarchitected_updateAnswerCmd.Flags().String("lens-alias", "", "")
	wellarchitected_updateAnswerCmd.Flags().String("notes", "", "")
	wellarchitected_updateAnswerCmd.Flags().String("question-id", "", "")
	wellarchitected_updateAnswerCmd.Flags().String("reason", "", "The reason why a question is not applicable to your workload.")
	wellarchitected_updateAnswerCmd.Flags().String("selected-choices", "", "")
	wellarchitected_updateAnswerCmd.Flags().String("workload-id", "", "")
	wellarchitected_updateAnswerCmd.MarkFlagRequired("lens-alias")
	wellarchitected_updateAnswerCmd.MarkFlagRequired("question-id")
	wellarchitected_updateAnswerCmd.MarkFlagRequired("workload-id")
	wellarchitectedCmd.AddCommand(wellarchitected_updateAnswerCmd)
}
