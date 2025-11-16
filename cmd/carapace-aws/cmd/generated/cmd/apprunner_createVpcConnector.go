package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_createVpcConnectorCmd = &cobra.Command{
	Use:   "create-vpc-connector",
	Short: "Create an App Runner VPC connector resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_createVpcConnectorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunner_createVpcConnectorCmd).Standalone()

		apprunner_createVpcConnectorCmd.Flags().String("security-groups", "", "A list of IDs of security groups that App Runner should use for access to Amazon Web Services resources under the specified subnets.")
		apprunner_createVpcConnectorCmd.Flags().String("subnets", "", "A list of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC.")
		apprunner_createVpcConnectorCmd.Flags().String("tags", "", "A list of metadata items that you can associate with your VPC connector resource.")
		apprunner_createVpcConnectorCmd.Flags().String("vpc-connector-name", "", "A name for the VPC connector.")
		apprunner_createVpcConnectorCmd.MarkFlagRequired("subnets")
		apprunner_createVpcConnectorCmd.MarkFlagRequired("vpc-connector-name")
	})
	apprunnerCmd.AddCommand(apprunner_createVpcConnectorCmd)
}
