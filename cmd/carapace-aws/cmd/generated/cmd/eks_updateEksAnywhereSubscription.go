package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_updateEksAnywhereSubscriptionCmd = &cobra.Command{
	Use:   "update-eks-anywhere-subscription",
	Short: "Update an EKS Anywhere Subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_updateEksAnywhereSubscriptionCmd).Standalone()

	eks_updateEksAnywhereSubscriptionCmd.Flags().Bool("auto-renew", false, "A boolean indicating whether or not to automatically renew the subscription.")
	eks_updateEksAnywhereSubscriptionCmd.Flags().String("client-request-token", "", "Unique, case-sensitive identifier to ensure the idempotency of the request.")
	eks_updateEksAnywhereSubscriptionCmd.Flags().String("id", "", "The ID of the subscription.")
	eks_updateEksAnywhereSubscriptionCmd.Flags().Bool("no-auto-renew", false, "A boolean indicating whether or not to automatically renew the subscription.")
	eks_updateEksAnywhereSubscriptionCmd.MarkFlagRequired("auto-renew")
	eks_updateEksAnywhereSubscriptionCmd.MarkFlagRequired("id")
	eks_updateEksAnywhereSubscriptionCmd.Flag("no-auto-renew").Hidden = true
	eks_updateEksAnywhereSubscriptionCmd.MarkFlagRequired("no-auto-renew")
	eksCmd.AddCommand(eks_updateEksAnywhereSubscriptionCmd)
}
