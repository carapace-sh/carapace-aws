package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_listDeploymentsCmd = &cobra.Command{
	Use:   "list-deployments",
	Short: "Returns a history of deployments for the group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_listDeploymentsCmd).Standalone()

	greengrass_listDeploymentsCmd.Flags().String("group-id", "", "The ID of the Greengrass group.")
	greengrass_listDeploymentsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	greengrass_listDeploymentsCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
	greengrass_listDeploymentsCmd.MarkFlagRequired("group-id")
	greengrassCmd.AddCommand(greengrass_listDeploymentsCmd)
}
