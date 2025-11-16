package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_deleteVpcIngressConnectionCmd = &cobra.Command{
	Use:   "delete-vpc-ingress-connection",
	Short: "Delete an App Runner VPC Ingress Connection resource that's associated with an App Runner service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_deleteVpcIngressConnectionCmd).Standalone()

	apprunner_deleteVpcIngressConnectionCmd.Flags().String("vpc-ingress-connection-arn", "", "The Amazon Resource Name (ARN) of the App Runner VPC Ingress Connection that you want to delete.")
	apprunner_deleteVpcIngressConnectionCmd.MarkFlagRequired("vpc-ingress-connection-arn")
	apprunnerCmd.AddCommand(apprunner_deleteVpcIngressConnectionCmd)
}
