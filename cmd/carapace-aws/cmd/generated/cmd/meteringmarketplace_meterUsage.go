package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var meteringmarketplace_meterUsageCmd = &cobra.Command{
	Use:   "meter-usage",
	Short: "API to emit metering records.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(meteringmarketplace_meterUsageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(meteringmarketplace_meterUsageCmd).Standalone()

		meteringmarketplace_meterUsageCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		meteringmarketplace_meterUsageCmd.Flags().Bool("dry-run", false, "Checks whether you have the permissions required for the action, but does not make the request.")
		meteringmarketplace_meterUsageCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the permissions required for the action, but does not make the request.")
		meteringmarketplace_meterUsageCmd.Flags().String("product-code", "", "Product code is used to uniquely identify a product in Amazon Web Services Marketplace.")
		meteringmarketplace_meterUsageCmd.Flags().String("timestamp", "", "Timestamp, in UTC, for which the usage is being reported.")
		meteringmarketplace_meterUsageCmd.Flags().String("usage-allocations", "", "The set of `UsageAllocations` to submit.")
		meteringmarketplace_meterUsageCmd.Flags().String("usage-dimension", "", "It will be one of the fcp dimension name provided during the publishing of the product.")
		meteringmarketplace_meterUsageCmd.Flags().String("usage-quantity", "", "Consumption value for the hour.")
		meteringmarketplace_meterUsageCmd.Flag("no-dry-run").Hidden = true
		meteringmarketplace_meterUsageCmd.MarkFlagRequired("product-code")
		meteringmarketplace_meterUsageCmd.MarkFlagRequired("timestamp")
		meteringmarketplace_meterUsageCmd.MarkFlagRequired("usage-dimension")
	})
	meteringmarketplaceCmd.AddCommand(meteringmarketplace_meterUsageCmd)
}
