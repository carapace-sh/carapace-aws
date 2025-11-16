package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_createStorageProfileCmd = &cobra.Command{
	Use:   "create-storage-profile",
	Short: "Creates a storage profile that specifies the operating system, file type, and file location of resources used on a farm.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_createStorageProfileCmd).Standalone()

	deadline_createStorageProfileCmd.Flags().String("client-token", "", "The unique token which the server uses to recognize retries of the same request.")
	deadline_createStorageProfileCmd.Flags().String("display-name", "", "The display name of the storage profile.")
	deadline_createStorageProfileCmd.Flags().String("farm-id", "", "The farm ID of the farm to connect to the storage profile.")
	deadline_createStorageProfileCmd.Flags().String("file-system-locations", "", "File system paths to include in the storage profile.")
	deadline_createStorageProfileCmd.Flags().String("os-family", "", "The type of operating system (OS) for the storage profile.")
	deadline_createStorageProfileCmd.MarkFlagRequired("display-name")
	deadline_createStorageProfileCmd.MarkFlagRequired("farm-id")
	deadline_createStorageProfileCmd.MarkFlagRequired("os-family")
	deadlineCmd.AddCommand(deadline_createStorageProfileCmd)
}
