package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_listDeploymentInstancesCmd = &cobra.Command{
	Use:   "list-deployment-instances",
	Short: "The newer `BatchGetDeploymentTargets` should be used instead because it works with all compute types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_listDeploymentInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_listDeploymentInstancesCmd).Standalone()

		codedeploy_listDeploymentInstancesCmd.Flags().String("deployment-id", "", "The unique ID of a deployment.")
		codedeploy_listDeploymentInstancesCmd.Flags().String("instance-status-filter", "", "A subset of instances to list by status:")
		codedeploy_listDeploymentInstancesCmd.Flags().String("instance-type-filter", "", "The set of instances in a blue/green deployment, either those in the original environment (\"BLUE\") or those in the replacement environment (\"GREEN\"), for which you want to view instance information.")
		codedeploy_listDeploymentInstancesCmd.Flags().String("next-token", "", "An identifier returned from the previous list deployment instances call.")
		codedeploy_listDeploymentInstancesCmd.MarkFlagRequired("deployment-id")
	})
	codedeployCmd.AddCommand(codedeploy_listDeploymentInstancesCmd)
}
