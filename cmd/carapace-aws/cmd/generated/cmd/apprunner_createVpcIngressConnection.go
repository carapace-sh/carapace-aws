package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_createVpcIngressConnectionCmd = &cobra.Command{
	Use:   "create-vpc-ingress-connection",
	Short: "Create an App Runner VPC Ingress Connection resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_createVpcIngressConnectionCmd).Standalone()

	apprunner_createVpcIngressConnectionCmd.Flags().String("ingress-vpc-configuration", "", "Specifications for the customer’s Amazon VPC and the related Amazon Web Services PrivateLink VPC endpoint that are used to create the VPC Ingress Connection resource.")
	apprunner_createVpcIngressConnectionCmd.Flags().String("service-arn", "", "The Amazon Resource Name (ARN) for this App Runner service that is used to create the VPC Ingress Connection resource.")
	apprunner_createVpcIngressConnectionCmd.Flags().String("tags", "", "An optional list of metadata items that you can associate with the VPC Ingress Connection resource.")
	apprunner_createVpcIngressConnectionCmd.Flags().String("vpc-ingress-connection-name", "", "A name for the VPC Ingress Connection resource.")
	apprunner_createVpcIngressConnectionCmd.MarkFlagRequired("ingress-vpc-configuration")
	apprunner_createVpcIngressConnectionCmd.MarkFlagRequired("service-arn")
	apprunner_createVpcIngressConnectionCmd.MarkFlagRequired("vpc-ingress-connection-name")
	apprunnerCmd.AddCommand(apprunner_createVpcIngressConnectionCmd)
}
