package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_describeVpcConnectorCmd = &cobra.Command{
	Use:   "describe-vpc-connector",
	Short: "Return a description of an App Runner VPC connector resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_describeVpcConnectorCmd).Standalone()

	apprunner_describeVpcConnectorCmd.Flags().String("vpc-connector-arn", "", "The Amazon Resource Name (ARN) of the App Runner VPC connector that you want a description for.")
	apprunner_describeVpcConnectorCmd.MarkFlagRequired("vpc-connector-arn")
	apprunnerCmd.AddCommand(apprunner_describeVpcConnectorCmd)
}
