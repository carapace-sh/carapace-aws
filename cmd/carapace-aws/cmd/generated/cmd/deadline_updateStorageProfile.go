package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_updateStorageProfileCmd = &cobra.Command{
	Use:   "update-storage-profile",
	Short: "Updates a storage profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_updateStorageProfileCmd).Standalone()

	deadline_updateStorageProfileCmd.Flags().String("client-token", "", "The unique token which the server uses to recognize retries of the same request.")
	deadline_updateStorageProfileCmd.Flags().String("display-name", "", "The display name of the storage profile to update.")
	deadline_updateStorageProfileCmd.Flags().String("farm-id", "", "The farm ID to update.")
	deadline_updateStorageProfileCmd.Flags().String("file-system-locations-to-add", "", "The file system location names to add.")
	deadline_updateStorageProfileCmd.Flags().String("file-system-locations-to-remove", "", "The file system location names to remove.")
	deadline_updateStorageProfileCmd.Flags().String("os-family", "", "The OS system to update.")
	deadline_updateStorageProfileCmd.Flags().String("storage-profile-id", "", "The storage profile ID to update.")
	deadline_updateStorageProfileCmd.MarkFlagRequired("farm-id")
	deadline_updateStorageProfileCmd.MarkFlagRequired("storage-profile-id")
	deadlineCmd.AddCommand(deadline_updateStorageProfileCmd)
}
