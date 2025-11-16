package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var launchWizard_listDeploymentEventsCmd = &cobra.Command{
	Use:   "list-deployment-events",
	Short: "Lists the events of a deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(launchWizard_listDeploymentEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(launchWizard_listDeploymentEventsCmd).Standalone()

		launchWizard_listDeploymentEventsCmd.Flags().String("deployment-id", "", "The ID of the deployment.")
		launchWizard_listDeploymentEventsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		launchWizard_listDeploymentEventsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		launchWizard_listDeploymentEventsCmd.MarkFlagRequired("deployment-id")
	})
	launchWizardCmd.AddCommand(launchWizard_listDeploymentEventsCmd)
}
