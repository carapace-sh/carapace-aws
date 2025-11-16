package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_notifyResourceDeploymentStatusChangeCmd = &cobra.Command{
	Use:   "notify-resource-deployment-status-change",
	Short: "Notify Proton of status changes to a provisioned resource when you use self-managed provisioning.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_notifyResourceDeploymentStatusChangeCmd).Standalone()

	proton_notifyResourceDeploymentStatusChangeCmd.Flags().String("deployment-id", "", "The deployment ID for your provisioned resource.")
	proton_notifyResourceDeploymentStatusChangeCmd.Flags().String("outputs", "", "The provisioned resource state change detail data that's returned by Proton.")
	proton_notifyResourceDeploymentStatusChangeCmd.Flags().String("resource-arn", "", "The provisioned resource Amazon Resource Name (ARN).")
	proton_notifyResourceDeploymentStatusChangeCmd.Flags().String("status", "", "The status of your provisioned resource.")
	proton_notifyResourceDeploymentStatusChangeCmd.Flags().String("status-message", "", "The deployment status message for your provisioned resource.")
	proton_notifyResourceDeploymentStatusChangeCmd.MarkFlagRequired("resource-arn")
	protonCmd.AddCommand(proton_notifyResourceDeploymentStatusChangeCmd)
}
