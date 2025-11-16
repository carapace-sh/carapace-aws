package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_listBulkDeploymentDetailedReportsCmd = &cobra.Command{
	Use:   "list-bulk-deployment-detailed-reports",
	Short: "Gets a paginated list of the deployments that have been started in a bulk deployment operation, and their current deployment status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_listBulkDeploymentDetailedReportsCmd).Standalone()

	greengrass_listBulkDeploymentDetailedReportsCmd.Flags().String("bulk-deployment-id", "", "The ID of the bulk deployment.")
	greengrass_listBulkDeploymentDetailedReportsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	greengrass_listBulkDeploymentDetailedReportsCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
	greengrass_listBulkDeploymentDetailedReportsCmd.MarkFlagRequired("bulk-deployment-id")
	greengrassCmd.AddCommand(greengrass_listBulkDeploymentDetailedReportsCmd)
}
