package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_updateAddonCmd = &cobra.Command{
	Use:   "update-addon",
	Short: "Updates an Amazon EKS add-on.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_updateAddonCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_updateAddonCmd).Standalone()

		eks_updateAddonCmd.Flags().String("addon-name", "", "The name of the add-on.")
		eks_updateAddonCmd.Flags().String("addon-version", "", "The version of the add-on.")
		eks_updateAddonCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		eks_updateAddonCmd.Flags().String("cluster-name", "", "The name of your cluster.")
		eks_updateAddonCmd.Flags().String("configuration-values", "", "The set of configuration values for the add-on that's created.")
		eks_updateAddonCmd.Flags().String("pod-identity-associations", "", "An array of EKS Pod Identity associations to be updated.")
		eks_updateAddonCmd.Flags().String("resolve-conflicts", "", "How to resolve field value conflicts for an Amazon EKS add-on if you've changed a value from the Amazon EKS default value.")
		eks_updateAddonCmd.Flags().String("service-account-role-arn", "", "The Amazon Resource Name (ARN) of an existing IAM role to bind to the add-on's service account.")
		eks_updateAddonCmd.MarkFlagRequired("addon-name")
		eks_updateAddonCmd.MarkFlagRequired("cluster-name")
	})
	eksCmd.AddCommand(eks_updateAddonCmd)
}
