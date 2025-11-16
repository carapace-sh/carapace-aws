package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_listAnomalousLogGroupsCmd = &cobra.Command{
	Use:   "list-anomalous-log-groups",
	Short: "Returns the list of log groups that contain log anomalies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_listAnomalousLogGroupsCmd).Standalone()

	devopsGuru_listAnomalousLogGroupsCmd.Flags().String("insight-id", "", "The ID of the insight containing the log groups.")
	devopsGuru_listAnomalousLogGroupsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	devopsGuru_listAnomalousLogGroupsCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
	devopsGuru_listAnomalousLogGroupsCmd.MarkFlagRequired("insight-id")
	devopsGuruCmd.AddCommand(devopsGuru_listAnomalousLogGroupsCmd)
}
