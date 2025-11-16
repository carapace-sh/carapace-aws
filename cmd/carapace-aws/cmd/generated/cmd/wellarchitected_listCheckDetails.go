package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_listCheckDetailsCmd = &cobra.Command{
	Use:   "list-check-details",
	Short: "List of Trusted Advisor check details by account related to the workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_listCheckDetailsCmd).Standalone()

	wellarchitected_listCheckDetailsCmd.Flags().String("choice-id", "", "")
	wellarchitected_listCheckDetailsCmd.Flags().String("lens-arn", "", "Well-Architected Lens ARN.")
	wellarchitected_listCheckDetailsCmd.Flags().String("max-results", "", "")
	wellarchitected_listCheckDetailsCmd.Flags().String("next-token", "", "")
	wellarchitected_listCheckDetailsCmd.Flags().String("pillar-id", "", "")
	wellarchitected_listCheckDetailsCmd.Flags().String("question-id", "", "")
	wellarchitected_listCheckDetailsCmd.Flags().String("workload-id", "", "")
	wellarchitected_listCheckDetailsCmd.MarkFlagRequired("choice-id")
	wellarchitected_listCheckDetailsCmd.MarkFlagRequired("lens-arn")
	wellarchitected_listCheckDetailsCmd.MarkFlagRequired("pillar-id")
	wellarchitected_listCheckDetailsCmd.MarkFlagRequired("question-id")
	wellarchitected_listCheckDetailsCmd.MarkFlagRequired("workload-id")
	wellarchitectedCmd.AddCommand(wellarchitected_listCheckDetailsCmd)
}
