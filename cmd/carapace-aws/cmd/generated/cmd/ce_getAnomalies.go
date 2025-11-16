package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_getAnomaliesCmd = &cobra.Command{
	Use:   "get-anomalies",
	Short: "Retrieves all of the cost anomalies detected on your account during the time period that's specified by the `DateInterval` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_getAnomaliesCmd).Standalone()

	ce_getAnomaliesCmd.Flags().String("date-interval", "", "Assigns the start and end dates for retrieving cost anomalies.")
	ce_getAnomaliesCmd.Flags().String("feedback", "", "Filters anomaly results by the feedback field on the anomaly object.")
	ce_getAnomaliesCmd.Flags().String("max-results", "", "The number of entries a paginated response contains.")
	ce_getAnomaliesCmd.Flags().String("monitor-arn", "", "Retrieves all of the cost anomalies detected for a specific cost anomaly monitor Amazon Resource Name (ARN).")
	ce_getAnomaliesCmd.Flags().String("next-page-token", "", "The token to retrieve the next set of results.")
	ce_getAnomaliesCmd.Flags().String("total-impact", "", "Filters anomaly results by the total impact field on the anomaly object.")
	ce_getAnomaliesCmd.MarkFlagRequired("date-interval")
	ceCmd.AddCommand(ce_getAnomaliesCmd)
}
