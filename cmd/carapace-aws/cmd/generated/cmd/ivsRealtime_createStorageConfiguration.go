package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_createStorageConfigurationCmd = &cobra.Command{
	Use:   "create-storage-configuration",
	Short: "Creates a new storage configuration, used to enable recording to Amazon S3.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_createStorageConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivsRealtime_createStorageConfigurationCmd).Standalone()

		ivsRealtime_createStorageConfigurationCmd.Flags().String("name", "", "Storage configuration name.")
		ivsRealtime_createStorageConfigurationCmd.Flags().String("s3", "", "A complex type that contains a storage configuration for where recorded video will be stored.")
		ivsRealtime_createStorageConfigurationCmd.Flags().String("tags", "", "Tags attached to the resource.")
		ivsRealtime_createStorageConfigurationCmd.MarkFlagRequired("s3")
	})
	ivsRealtimeCmd.AddCommand(ivsRealtime_createStorageConfigurationCmd)
}
