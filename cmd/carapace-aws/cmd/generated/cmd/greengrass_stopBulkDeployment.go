package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_stopBulkDeploymentCmd = &cobra.Command{
	Use:   "stop-bulk-deployment",
	Short: "Stops the execution of a bulk deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_stopBulkDeploymentCmd).Standalone()

	greengrass_stopBulkDeploymentCmd.Flags().String("bulk-deployment-id", "", "The ID of the bulk deployment.")
	greengrass_stopBulkDeploymentCmd.MarkFlagRequired("bulk-deployment-id")
	greengrassCmd.AddCommand(greengrass_stopBulkDeploymentCmd)
}
