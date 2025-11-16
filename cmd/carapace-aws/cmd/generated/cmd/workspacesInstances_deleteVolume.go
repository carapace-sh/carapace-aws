package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesInstances_deleteVolumeCmd = &cobra.Command{
	Use:   "delete-volume",
	Short: "Deletes a specified volume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesInstances_deleteVolumeCmd).Standalone()

	workspacesInstances_deleteVolumeCmd.Flags().String("volume-id", "", "Identifier of the volume to delete.")
	workspacesInstances_deleteVolumeCmd.MarkFlagRequired("volume-id")
	workspacesInstancesCmd.AddCommand(workspacesInstances_deleteVolumeCmd)
}
