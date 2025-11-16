package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_listDeploymentsCmd = &cobra.Command{
	Use:   "list-deployments",
	Short: "Lists the deployments for an environment in descending deployment number order.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_listDeploymentsCmd).Standalone()

	appconfig_listDeploymentsCmd.Flags().String("application-id", "", "The application ID.")
	appconfig_listDeploymentsCmd.Flags().String("environment-id", "", "The environment ID.")
	appconfig_listDeploymentsCmd.Flags().String("max-results", "", "The maximum number of items that may be returned for this call.")
	appconfig_listDeploymentsCmd.Flags().String("next-token", "", "The token returned by a prior call to this operation indicating the next set of results to be returned.")
	appconfig_listDeploymentsCmd.MarkFlagRequired("application-id")
	appconfig_listDeploymentsCmd.MarkFlagRequired("environment-id")
	appconfigCmd.AddCommand(appconfig_listDeploymentsCmd)
}
