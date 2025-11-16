package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_createAddonCmd = &cobra.Command{
	Use:   "create-addon",
	Short: "Creates an Amazon EKS add-on.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_createAddonCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_createAddonCmd).Standalone()

		eks_createAddonCmd.Flags().String("addon-name", "", "The name of the add-on.")
		eks_createAddonCmd.Flags().String("addon-version", "", "The version of the add-on.")
		eks_createAddonCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		eks_createAddonCmd.Flags().String("cluster-name", "", "The name of your cluster.")
		eks_createAddonCmd.Flags().String("configuration-values", "", "The set of configuration values for the add-on that's created.")
		eks_createAddonCmd.Flags().String("namespace-config", "", "The namespace configuration for the addon.")
		eks_createAddonCmd.Flags().String("pod-identity-associations", "", "An array of EKS Pod Identity associations to be created.")
		eks_createAddonCmd.Flags().String("resolve-conflicts", "", "How to resolve field value conflicts for an Amazon EKS add-on.")
		eks_createAddonCmd.Flags().String("service-account-role-arn", "", "The Amazon Resource Name (ARN) of an existing IAM role to bind to the add-on's service account.")
		eks_createAddonCmd.Flags().String("tags", "", "Metadata that assists with categorization and organization.")
		eks_createAddonCmd.MarkFlagRequired("addon-name")
		eks_createAddonCmd.MarkFlagRequired("cluster-name")
	})
	eksCmd.AddCommand(eks_createAddonCmd)
}
