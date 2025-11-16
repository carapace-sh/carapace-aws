package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_deleteKxVolumeCmd = &cobra.Command{
	Use:   "delete-kx-volume",
	Short: "Deletes a volume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_deleteKxVolumeCmd).Standalone()

	finspace_deleteKxVolumeCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
	finspace_deleteKxVolumeCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment, whose clusters can attach to the volume.")
	finspace_deleteKxVolumeCmd.Flags().String("volume-name", "", "The name of the volume that you want to delete.")
	finspace_deleteKxVolumeCmd.MarkFlagRequired("environment-id")
	finspace_deleteKxVolumeCmd.MarkFlagRequired("volume-name")
	finspaceCmd.AddCommand(finspace_deleteKxVolumeCmd)
}
