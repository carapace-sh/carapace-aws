package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_deleteStorageConfigurationCmd = &cobra.Command{
	Use:   "delete-storage-configuration",
	Short: "Deletes the storage configuration for the specified ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_deleteStorageConfigurationCmd).Standalone()

	ivsRealtime_deleteStorageConfigurationCmd.Flags().String("arn", "", "ARN of the storage configuration to be deleted.")
	ivsRealtime_deleteStorageConfigurationCmd.MarkFlagRequired("arn")
	ivsRealtimeCmd.AddCommand(ivsRealtime_deleteStorageConfigurationCmd)
}
