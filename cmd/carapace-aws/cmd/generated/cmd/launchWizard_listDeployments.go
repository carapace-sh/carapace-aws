package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var launchWizard_listDeploymentsCmd = &cobra.Command{
	Use:   "list-deployments",
	Short: "Lists the deployments that have been created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(launchWizard_listDeploymentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(launchWizard_listDeploymentsCmd).Standalone()

		launchWizard_listDeploymentsCmd.Flags().String("filters", "", "Filters to scope the results.")
		launchWizard_listDeploymentsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		launchWizard_listDeploymentsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	})
	launchWizardCmd.AddCommand(launchWizard_listDeploymentsCmd)
}
