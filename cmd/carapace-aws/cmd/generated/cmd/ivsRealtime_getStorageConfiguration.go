package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_getStorageConfigurationCmd = &cobra.Command{
	Use:   "get-storage-configuration",
	Short: "Gets the storage configuration for the specified ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_getStorageConfigurationCmd).Standalone()

	ivsRealtime_getStorageConfigurationCmd.Flags().String("arn", "", "ARN of the storage configuration to be retrieved.")
	ivsRealtime_getStorageConfigurationCmd.MarkFlagRequired("arn")
	ivsRealtimeCmd.AddCommand(ivsRealtime_getStorageConfigurationCmd)
}
