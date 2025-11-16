package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_cancelComponentDeploymentCmd = &cobra.Command{
	Use:   "cancel-component-deployment",
	Short: "Attempts to cancel a component deployment (for a component that is in the `IN_PROGRESS` deployment status).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_cancelComponentDeploymentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_cancelComponentDeploymentCmd).Standalone()

		proton_cancelComponentDeploymentCmd.Flags().String("component-name", "", "The name of the component with the deployment to cancel.")
		proton_cancelComponentDeploymentCmd.MarkFlagRequired("component-name")
	})
	protonCmd.AddCommand(proton_cancelComponentDeploymentCmd)
}
