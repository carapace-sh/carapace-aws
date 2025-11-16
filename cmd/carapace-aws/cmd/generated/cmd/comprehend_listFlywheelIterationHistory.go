package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_listFlywheelIterationHistoryCmd = &cobra.Command{
	Use:   "list-flywheel-iteration-history",
	Short: "Information about the history of a flywheel iteration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_listFlywheelIterationHistoryCmd).Standalone()

	comprehend_listFlywheelIterationHistoryCmd.Flags().String("filter", "", "Filter the flywheel iteration history based on creation time.")
	comprehend_listFlywheelIterationHistoryCmd.Flags().String("flywheel-arn", "", "The ARN of the flywheel.")
	comprehend_listFlywheelIterationHistoryCmd.Flags().String("max-results", "", "Maximum number of iteration history results to return")
	comprehend_listFlywheelIterationHistoryCmd.Flags().String("next-token", "", "Next token")
	comprehend_listFlywheelIterationHistoryCmd.MarkFlagRequired("flywheel-arn")
	comprehendCmd.AddCommand(comprehend_listFlywheelIterationHistoryCmd)
}
