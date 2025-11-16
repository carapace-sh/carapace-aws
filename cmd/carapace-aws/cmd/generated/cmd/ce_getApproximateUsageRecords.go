package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_getApproximateUsageRecordsCmd = &cobra.Command{
	Use:   "get-approximate-usage-records",
	Short: "Retrieves estimated usage records for hourly granularity or resource-level data at daily granularity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_getApproximateUsageRecordsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ce_getApproximateUsageRecordsCmd).Standalone()

		ce_getApproximateUsageRecordsCmd.Flags().String("approximation-dimension", "", "The service to evaluate for the usage records.")
		ce_getApproximateUsageRecordsCmd.Flags().String("granularity", "", "How granular you want the data to be.")
		ce_getApproximateUsageRecordsCmd.Flags().String("services", "", "The service metadata for the service or services you want to query.")
		ce_getApproximateUsageRecordsCmd.MarkFlagRequired("approximation-dimension")
		ce_getApproximateUsageRecordsCmd.MarkFlagRequired("granularity")
	})
	ceCmd.AddCommand(ce_getApproximateUsageRecordsCmd)
}
