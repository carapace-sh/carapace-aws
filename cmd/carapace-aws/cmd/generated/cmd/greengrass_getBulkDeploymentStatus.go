package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getBulkDeploymentStatusCmd = &cobra.Command{
	Use:   "get-bulk-deployment-status",
	Short: "Returns the status of a bulk deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getBulkDeploymentStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_getBulkDeploymentStatusCmd).Standalone()

		greengrass_getBulkDeploymentStatusCmd.Flags().String("bulk-deployment-id", "", "The ID of the bulk deployment.")
		greengrass_getBulkDeploymentStatusCmd.MarkFlagRequired("bulk-deployment-id")
	})
	greengrassCmd.AddCommand(greengrass_getBulkDeploymentStatusCmd)
}
