package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var meteringmarketplace_registerUsageCmd = &cobra.Command{
	Use:   "register-usage",
	Short: "Paid container software products sold through Amazon Web Services Marketplace must integrate with the Amazon Web Services Marketplace Metering Service and call the `RegisterUsage` operation for software entitlement and metering.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(meteringmarketplace_registerUsageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(meteringmarketplace_registerUsageCmd).Standalone()

		meteringmarketplace_registerUsageCmd.Flags().String("nonce", "", "(Optional) To scope down the registration to a specific running software instance and guard against replay attacks.")
		meteringmarketplace_registerUsageCmd.Flags().String("product-code", "", "Product code is used to uniquely identify a product in Amazon Web Services Marketplace.")
		meteringmarketplace_registerUsageCmd.Flags().String("public-key-version", "", "Public Key Version provided by Amazon Web Services Marketplace")
		meteringmarketplace_registerUsageCmd.MarkFlagRequired("product-code")
		meteringmarketplace_registerUsageCmd.MarkFlagRequired("public-key-version")
	})
	meteringmarketplaceCmd.AddCommand(meteringmarketplace_registerUsageCmd)
}
