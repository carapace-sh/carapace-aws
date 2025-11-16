package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_deleteStorageProfileCmd = &cobra.Command{
	Use:   "delete-storage-profile",
	Short: "Deletes a storage profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_deleteStorageProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_deleteStorageProfileCmd).Standalone()

		deadline_deleteStorageProfileCmd.Flags().String("farm-id", "", "The farm ID of the farm from which to remove the storage profile.")
		deadline_deleteStorageProfileCmd.Flags().String("storage-profile-id", "", "The storage profile ID of the storage profile to delete.")
		deadline_deleteStorageProfileCmd.MarkFlagRequired("farm-id")
		deadline_deleteStorageProfileCmd.MarkFlagRequired("storage-profile-id")
	})
	deadlineCmd.AddCommand(deadline_deleteStorageProfileCmd)
}
