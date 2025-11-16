package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_updateProblemCmd = &cobra.Command{
	Use:   "update-problem",
	Short: "Updates the visibility of the problem or specifies the problem as `RESOLVED`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_updateProblemCmd).Standalone()

	applicationInsights_updateProblemCmd.Flags().String("problem-id", "", "The ID of the problem.")
	applicationInsights_updateProblemCmd.Flags().String("update-status", "", "The status of the problem.")
	applicationInsights_updateProblemCmd.Flags().String("visibility", "", "The visibility of a problem.")
	applicationInsights_updateProblemCmd.MarkFlagRequired("problem-id")
	applicationInsightsCmd.AddCommand(applicationInsights_updateProblemCmd)
}
