package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_getDeploymentCmd = &cobra.Command{
	Use:   "get-deployment",
	Short: "Gets information about a deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_getDeploymentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_getDeploymentCmd).Standalone()

		codedeploy_getDeploymentCmd.Flags().String("deployment-id", "", "The unique ID of a deployment associated with the user or Amazon Web Services account.")
		codedeploy_getDeploymentCmd.MarkFlagRequired("deployment-id")
	})
	codedeployCmd.AddCommand(codedeploy_getDeploymentCmd)
}
