package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendraRanking_createRescoreExecutionPlanCmd = &cobra.Command{
	Use:   "create-rescore-execution-plan",
	Short: "Creates a rescore execution plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendraRanking_createRescoreExecutionPlanCmd).Standalone()

	kendraRanking_createRescoreExecutionPlanCmd.Flags().String("capacity-units", "", "You can set additional capacity units to meet the needs of your rescore execution plan.")
	kendraRanking_createRescoreExecutionPlanCmd.Flags().String("client-token", "", "A token that you provide to identify the request to create a rescore execution plan.")
	kendraRanking_createRescoreExecutionPlanCmd.Flags().String("description", "", "A description for the rescore execution plan.")
	kendraRanking_createRescoreExecutionPlanCmd.Flags().String("name", "", "A name for the rescore execution plan.")
	kendraRanking_createRescoreExecutionPlanCmd.Flags().String("tags", "", "A list of key-value pairs that identify or categorize your rescore execution plan.")
	kendraRanking_createRescoreExecutionPlanCmd.MarkFlagRequired("name")
	kendraRankingCmd.AddCommand(kendraRanking_createRescoreExecutionPlanCmd)
}
