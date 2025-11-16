package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_updateVpcIngressConnectionCmd = &cobra.Command{
	Use:   "update-vpc-ingress-connection",
	Short: "Update an existing App Runner VPC Ingress Connection resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_updateVpcIngressConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunner_updateVpcIngressConnectionCmd).Standalone()

		apprunner_updateVpcIngressConnectionCmd.Flags().String("ingress-vpc-configuration", "", "Specifications for the customer’s Amazon VPC and the related Amazon Web Services PrivateLink VPC endpoint that are used to update the VPC Ingress Connection resource.")
		apprunner_updateVpcIngressConnectionCmd.Flags().String("vpc-ingress-connection-arn", "", "The Amazon Resource Name (Arn) for the App Runner VPC Ingress Connection resource that you want to update.")
		apprunner_updateVpcIngressConnectionCmd.MarkFlagRequired("ingress-vpc-configuration")
		apprunner_updateVpcIngressConnectionCmd.MarkFlagRequired("vpc-ingress-connection-arn")
	})
	apprunnerCmd.AddCommand(apprunner_updateVpcIngressConnectionCmd)
}
