package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var meteringmarketplace_batchMeterUsageCmd = &cobra.Command{
	Use:   "batch-meter-usage",
	Short: "The `CustomerIdentifier` parameter is scheduled for deprecation on March 31, 2026. Use `CustomerAWSAccountID` instead.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(meteringmarketplace_batchMeterUsageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(meteringmarketplace_batchMeterUsageCmd).Standalone()

		meteringmarketplace_batchMeterUsageCmd.Flags().String("product-code", "", "Product code is used to uniquely identify a product in Amazon Web Services Marketplace.")
		meteringmarketplace_batchMeterUsageCmd.Flags().String("usage-records", "", "The set of `UsageRecords` to submit.")
		meteringmarketplace_batchMeterUsageCmd.MarkFlagRequired("product-code")
		meteringmarketplace_batchMeterUsageCmd.MarkFlagRequired("usage-records")
	})
	meteringmarketplaceCmd.AddCommand(meteringmarketplace_batchMeterUsageCmd)
}
