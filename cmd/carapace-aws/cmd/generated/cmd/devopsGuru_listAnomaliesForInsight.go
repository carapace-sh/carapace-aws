package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_listAnomaliesForInsightCmd = &cobra.Command{
	Use:   "list-anomalies-for-insight",
	Short: "Returns a list of the anomalies that belong to an insight that you specify using its ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_listAnomaliesForInsightCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devopsGuru_listAnomaliesForInsightCmd).Standalone()

		devopsGuru_listAnomaliesForInsightCmd.Flags().String("account-id", "", "The ID of the Amazon Web Services account.")
		devopsGuru_listAnomaliesForInsightCmd.Flags().String("filters", "", "Specifies one or more service names that are used to list anomalies.")
		devopsGuru_listAnomaliesForInsightCmd.Flags().String("insight-id", "", "The ID of the insight.")
		devopsGuru_listAnomaliesForInsightCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		devopsGuru_listAnomaliesForInsightCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
		devopsGuru_listAnomaliesForInsightCmd.Flags().String("start-time-range", "", "A time range used to specify when the requested anomalies started.")
		devopsGuru_listAnomaliesForInsightCmd.MarkFlagRequired("insight-id")
	})
	devopsGuruCmd.AddCommand(devopsGuru_listAnomaliesForInsightCmd)
}
