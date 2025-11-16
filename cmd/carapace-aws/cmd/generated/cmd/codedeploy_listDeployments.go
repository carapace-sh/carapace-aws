package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_listDeploymentsCmd = &cobra.Command{
	Use:   "list-deployments",
	Short: "Lists the deployments in a deployment group for an application registered with the user or Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_listDeploymentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_listDeploymentsCmd).Standalone()

		codedeploy_listDeploymentsCmd.Flags().String("application-name", "", "The name of an CodeDeploy application associated with the user or Amazon Web Services account.")
		codedeploy_listDeploymentsCmd.Flags().String("create-time-range", "", "A time range (start and end) for returning a subset of the list of deployments.")
		codedeploy_listDeploymentsCmd.Flags().String("deployment-group-name", "", "The name of a deployment group for the specified application.")
		codedeploy_listDeploymentsCmd.Flags().String("external-id", "", "The unique ID of an external resource for returning deployments linked to the external resource.")
		codedeploy_listDeploymentsCmd.Flags().String("include-only-statuses", "", "A subset of deployments to list by status:")
		codedeploy_listDeploymentsCmd.Flags().String("next-token", "", "An identifier returned from the previous list deployments call.")
	})
	codedeployCmd.AddCommand(codedeploy_listDeploymentsCmd)
}
