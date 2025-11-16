package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesInstances_disassociateVolumeCmd = &cobra.Command{
	Use:   "disassociate-volume",
	Short: "Detaches a volume from a WorkSpace Instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesInstances_disassociateVolumeCmd).Standalone()

	workspacesInstances_disassociateVolumeCmd.Flags().String("device", "", "Device path of volume to detach.")
	workspacesInstances_disassociateVolumeCmd.Flags().String("disassociate-mode", "", "Mode for volume detachment.")
	workspacesInstances_disassociateVolumeCmd.Flags().String("volume-id", "", "Volume to be detached.")
	workspacesInstances_disassociateVolumeCmd.Flags().String("workspace-instance-id", "", "WorkSpace Instance to detach volume from.")
	workspacesInstances_disassociateVolumeCmd.MarkFlagRequired("volume-id")
	workspacesInstances_disassociateVolumeCmd.MarkFlagRequired("workspace-instance-id")
	workspacesInstancesCmd.AddCommand(workspacesInstances_disassociateVolumeCmd)
}
