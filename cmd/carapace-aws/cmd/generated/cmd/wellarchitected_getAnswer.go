package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_getAnswerCmd = &cobra.Command{
	Use:   "get-answer",
	Short: "Get the answer to a specific question in a workload review.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_getAnswerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_getAnswerCmd).Standalone()

		wellarchitected_getAnswerCmd.Flags().String("lens-alias", "", "")
		wellarchitected_getAnswerCmd.Flags().String("milestone-number", "", "")
		wellarchitected_getAnswerCmd.Flags().String("question-id", "", "")
		wellarchitected_getAnswerCmd.Flags().String("workload-id", "", "")
		wellarchitected_getAnswerCmd.MarkFlagRequired("lens-alias")
		wellarchitected_getAnswerCmd.MarkFlagRequired("question-id")
		wellarchitected_getAnswerCmd.MarkFlagRequired("workload-id")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_getAnswerCmd)
}
