package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_listBulkDeploymentsCmd = &cobra.Command{
	Use:   "list-bulk-deployments",
	Short: "Returns a list of bulk deployments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_listBulkDeploymentsCmd).Standalone()

	greengrass_listBulkDeploymentsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	greengrass_listBulkDeploymentsCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
	greengrassCmd.AddCommand(greengrass_listBulkDeploymentsCmd)
}
