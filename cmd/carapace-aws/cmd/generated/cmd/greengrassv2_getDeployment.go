package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_getDeploymentCmd = &cobra.Command{
	Use:   "get-deployment",
	Short: "Gets a deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_getDeploymentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrassv2_getDeploymentCmd).Standalone()

		greengrassv2_getDeploymentCmd.Flags().String("deployment-id", "", "The ID of the deployment.")
		greengrassv2_getDeploymentCmd.MarkFlagRequired("deployment-id")
	})
	greengrassv2Cmd.AddCommand(greengrassv2_getDeploymentCmd)
}
