package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_describeVpcIngressConnectionCmd = &cobra.Command{
	Use:   "describe-vpc-ingress-connection",
	Short: "Return a full description of an App Runner VPC Ingress Connection resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_describeVpcIngressConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunner_describeVpcIngressConnectionCmd).Standalone()

		apprunner_describeVpcIngressConnectionCmd.Flags().String("vpc-ingress-connection-arn", "", "The Amazon Resource Name (ARN) of the App Runner VPC Ingress Connection that you want a description for.")
		apprunner_describeVpcIngressConnectionCmd.MarkFlagRequired("vpc-ingress-connection-arn")
	})
	apprunnerCmd.AddCommand(apprunner_describeVpcIngressConnectionCmd)
}
