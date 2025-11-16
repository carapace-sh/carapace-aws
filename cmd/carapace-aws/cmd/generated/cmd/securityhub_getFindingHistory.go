package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_getFindingHistoryCmd = &cobra.Command{
	Use:   "get-finding-history",
	Short: "Returns the history of a Security Hub finding.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_getFindingHistoryCmd).Standalone()

	securityhub_getFindingHistoryCmd.Flags().String("end-time", "", "An ISO 8601-formatted timestamp that indicates the end time of the requested finding history.")
	securityhub_getFindingHistoryCmd.Flags().String("finding-identifier", "", "")
	securityhub_getFindingHistoryCmd.Flags().String("max-results", "", "The maximum number of results to be returned.")
	securityhub_getFindingHistoryCmd.Flags().String("next-token", "", "A token for pagination purposes.")
	securityhub_getFindingHistoryCmd.Flags().String("start-time", "", "A timestamp that indicates the start time of the requested finding history.")
	securityhub_getFindingHistoryCmd.MarkFlagRequired("finding-identifier")
	securityhubCmd.AddCommand(securityhub_getFindingHistoryCmd)
}
