package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceDeployment_putDeploymentParameterCmd = &cobra.Command{
	Use:   "put-deployment-parameter",
	Short: "Creates or updates a deployment parameter and is targeted by `catalog` and `agreementId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceDeployment_putDeploymentParameterCmd).Standalone()

	marketplaceDeployment_putDeploymentParameterCmd.Flags().String("agreement-id", "", "The unique identifier of the agreement.")
	marketplaceDeployment_putDeploymentParameterCmd.Flags().String("catalog", "", "The catalog related to the request.")
	marketplaceDeployment_putDeploymentParameterCmd.Flags().String("client-token", "", "The idempotency token for deployment parameters.")
	marketplaceDeployment_putDeploymentParameterCmd.Flags().String("deployment-parameter", "", "The deployment parameter targeted to the acceptor of an agreement for which to create the AWS Secret Manager resource.")
	marketplaceDeployment_putDeploymentParameterCmd.Flags().String("expiration-date", "", "The date when deployment parameters expire and are scheduled for deletion.")
	marketplaceDeployment_putDeploymentParameterCmd.Flags().String("product-id", "", "The product for which AWS Marketplace will save secrets for the buyer’s account.")
	marketplaceDeployment_putDeploymentParameterCmd.Flags().String("tags", "", "A map of key-value pairs, where each pair represents a tag saved to the resource.")
	marketplaceDeployment_putDeploymentParameterCmd.MarkFlagRequired("agreement-id")
	marketplaceDeployment_putDeploymentParameterCmd.MarkFlagRequired("catalog")
	marketplaceDeployment_putDeploymentParameterCmd.MarkFlagRequired("deployment-parameter")
	marketplaceDeployment_putDeploymentParameterCmd.MarkFlagRequired("product-id")
	marketplaceDeploymentCmd.AddCommand(marketplaceDeployment_putDeploymentParameterCmd)
}
