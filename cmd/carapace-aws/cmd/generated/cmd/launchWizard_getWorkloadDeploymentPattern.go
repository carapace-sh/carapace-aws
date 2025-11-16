package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var launchWizard_getWorkloadDeploymentPatternCmd = &cobra.Command{
	Use:   "get-workload-deployment-pattern",
	Short: "Returns details for a given workload and deployment pattern, including the available specifications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(launchWizard_getWorkloadDeploymentPatternCmd).Standalone()

	launchWizard_getWorkloadDeploymentPatternCmd.Flags().String("deployment-pattern-name", "", "The name of the deployment pattern.")
	launchWizard_getWorkloadDeploymentPatternCmd.Flags().String("workload-name", "", "The name of the workload.")
	launchWizard_getWorkloadDeploymentPatternCmd.MarkFlagRequired("deployment-pattern-name")
	launchWizard_getWorkloadDeploymentPatternCmd.MarkFlagRequired("workload-name")
	launchWizardCmd.AddCommand(launchWizard_getWorkloadDeploymentPatternCmd)
}
