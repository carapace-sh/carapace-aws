package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_deleteVpcConnectorCmd = &cobra.Command{
	Use:   "delete-vpc-connector",
	Short: "Delete an App Runner VPC connector resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_deleteVpcConnectorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunner_deleteVpcConnectorCmd).Standalone()

		apprunner_deleteVpcConnectorCmd.Flags().String("vpc-connector-arn", "", "The Amazon Resource Name (ARN) of the App Runner VPC connector that you want to delete.")
		apprunner_deleteVpcConnectorCmd.MarkFlagRequired("vpc-connector-arn")
	})
	apprunnerCmd.AddCommand(apprunner_deleteVpcConnectorCmd)
}
