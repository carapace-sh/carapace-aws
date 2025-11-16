package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_getStorageProfileCmd = &cobra.Command{
	Use:   "get-storage-profile",
	Short: "Gets a storage profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_getStorageProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_getStorageProfileCmd).Standalone()

		deadline_getStorageProfileCmd.Flags().String("farm-id", "", "The farm ID for the storage profile.")
		deadline_getStorageProfileCmd.Flags().String("storage-profile-id", "", "The storage profile ID.")
		deadline_getStorageProfileCmd.MarkFlagRequired("farm-id")
		deadline_getStorageProfileCmd.MarkFlagRequired("storage-profile-id")
	})
	deadlineCmd.AddCommand(deadline_getStorageProfileCmd)
}
