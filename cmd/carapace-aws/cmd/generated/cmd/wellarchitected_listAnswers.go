package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_listAnswersCmd = &cobra.Command{
	Use:   "list-answers",
	Short: "List of answers for a particular workload and lens.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_listAnswersCmd).Standalone()

	wellarchitected_listAnswersCmd.Flags().String("lens-alias", "", "")
	wellarchitected_listAnswersCmd.Flags().String("max-results", "", "The maximum number of results to return for this request.")
	wellarchitected_listAnswersCmd.Flags().String("milestone-number", "", "")
	wellarchitected_listAnswersCmd.Flags().String("next-token", "", "")
	wellarchitected_listAnswersCmd.Flags().String("pillar-id", "", "")
	wellarchitected_listAnswersCmd.Flags().String("question-priority", "", "The priority of the question.")
	wellarchitected_listAnswersCmd.Flags().String("workload-id", "", "")
	wellarchitected_listAnswersCmd.MarkFlagRequired("lens-alias")
	wellarchitected_listAnswersCmd.MarkFlagRequired("workload-id")
	wellarchitectedCmd.AddCommand(wellarchitected_listAnswersCmd)
}
