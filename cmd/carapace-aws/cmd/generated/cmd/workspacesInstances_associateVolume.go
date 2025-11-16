package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesInstances_associateVolumeCmd = &cobra.Command{
	Use:   "associate-volume",
	Short: "Attaches a volume to a WorkSpace Instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesInstances_associateVolumeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesInstances_associateVolumeCmd).Standalone()

		workspacesInstances_associateVolumeCmd.Flags().String("device", "", "Device path for volume attachment.")
		workspacesInstances_associateVolumeCmd.Flags().String("volume-id", "", "Volume to be attached.")
		workspacesInstances_associateVolumeCmd.Flags().String("workspace-instance-id", "", "WorkSpace Instance to attach volume to.")
		workspacesInstances_associateVolumeCmd.MarkFlagRequired("device")
		workspacesInstances_associateVolumeCmd.MarkFlagRequired("volume-id")
		workspacesInstances_associateVolumeCmd.MarkFlagRequired("workspace-instance-id")
	})
	workspacesInstancesCmd.AddCommand(workspacesInstances_associateVolumeCmd)
}
