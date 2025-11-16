package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var launchWizard_listWorkloadDeploymentPatternsCmd = &cobra.Command{
	Use:   "list-workload-deployment-patterns",
	Short: "Lists the workload deployment patterns for a given workload name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(launchWizard_listWorkloadDeploymentPatternsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(launchWizard_listWorkloadDeploymentPatternsCmd).Standalone()

		launchWizard_listWorkloadDeploymentPatternsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		launchWizard_listWorkloadDeploymentPatternsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		launchWizard_listWorkloadDeploymentPatternsCmd.Flags().String("workload-name", "", "The name of the workload.")
		launchWizard_listWorkloadDeploymentPatternsCmd.MarkFlagRequired("workload-name")
	})
	launchWizardCmd.AddCommand(launchWizard_listWorkloadDeploymentPatternsCmd)
}
