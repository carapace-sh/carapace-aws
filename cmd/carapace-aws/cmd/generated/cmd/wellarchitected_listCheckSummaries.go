package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_listCheckSummariesCmd = &cobra.Command{
	Use:   "list-check-summaries",
	Short: "List of Trusted Advisor checks summarized for all accounts related to the workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_listCheckSummariesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_listCheckSummariesCmd).Standalone()

		wellarchitected_listCheckSummariesCmd.Flags().String("choice-id", "", "")
		wellarchitected_listCheckSummariesCmd.Flags().String("lens-arn", "", "Well-Architected Lens ARN.")
		wellarchitected_listCheckSummariesCmd.Flags().String("max-results", "", "")
		wellarchitected_listCheckSummariesCmd.Flags().String("next-token", "", "")
		wellarchitected_listCheckSummariesCmd.Flags().String("pillar-id", "", "")
		wellarchitected_listCheckSummariesCmd.Flags().String("question-id", "", "")
		wellarchitected_listCheckSummariesCmd.Flags().String("workload-id", "", "")
		wellarchitected_listCheckSummariesCmd.MarkFlagRequired("choice-id")
		wellarchitected_listCheckSummariesCmd.MarkFlagRequired("lens-arn")
		wellarchitected_listCheckSummariesCmd.MarkFlagRequired("pillar-id")
		wellarchitected_listCheckSummariesCmd.MarkFlagRequired("question-id")
		wellarchitected_listCheckSummariesCmd.MarkFlagRequired("workload-id")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_listCheckSummariesCmd)
}
