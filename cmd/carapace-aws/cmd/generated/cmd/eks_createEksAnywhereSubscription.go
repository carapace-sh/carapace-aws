package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_createEksAnywhereSubscriptionCmd = &cobra.Command{
	Use:   "create-eks-anywhere-subscription",
	Short: "Creates an EKS Anywhere subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_createEksAnywhereSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_createEksAnywhereSubscriptionCmd).Standalone()

		eks_createEksAnywhereSubscriptionCmd.Flags().Bool("auto-renew", false, "A boolean indicating whether the subscription auto renews at the end of the term.")
		eks_createEksAnywhereSubscriptionCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		eks_createEksAnywhereSubscriptionCmd.Flags().String("license-quantity", "", "The number of licenses to purchase with the subscription.")
		eks_createEksAnywhereSubscriptionCmd.Flags().String("license-type", "", "The license type for all licenses in the subscription.")
		eks_createEksAnywhereSubscriptionCmd.Flags().String("name", "", "The unique name for your subscription.")
		eks_createEksAnywhereSubscriptionCmd.Flags().Bool("no-auto-renew", false, "A boolean indicating whether the subscription auto renews at the end of the term.")
		eks_createEksAnywhereSubscriptionCmd.Flags().String("tags", "", "The metadata for a subscription to assist with categorization and organization.")
		eks_createEksAnywhereSubscriptionCmd.Flags().String("term", "", "An object representing the term duration and term unit type of your subscription.")
		eks_createEksAnywhereSubscriptionCmd.MarkFlagRequired("name")
		eks_createEksAnywhereSubscriptionCmd.Flag("no-auto-renew").Hidden = true
		eks_createEksAnywhereSubscriptionCmd.MarkFlagRequired("term")
	})
	eksCmd.AddCommand(eks_createEksAnywhereSubscriptionCmd)
}
