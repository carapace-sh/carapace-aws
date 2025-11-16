package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateInstanceStorageConfigCmd = &cobra.Command{
	Use:   "update-instance-storage-config",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateInstanceStorageConfigCmd).Standalone()

	connect_updateInstanceStorageConfigCmd.Flags().String("association-id", "", "The existing association identifier that uniquely identifies the resource type and storage config for the given instance ID.")
	connect_updateInstanceStorageConfigCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_updateInstanceStorageConfigCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updateInstanceStorageConfigCmd.Flags().String("resource-type", "", "A valid resource type.")
	connect_updateInstanceStorageConfigCmd.Flags().String("storage-config", "", "")
	connect_updateInstanceStorageConfigCmd.MarkFlagRequired("association-id")
	connect_updateInstanceStorageConfigCmd.MarkFlagRequired("instance-id")
	connect_updateInstanceStorageConfigCmd.MarkFlagRequired("resource-type")
	connect_updateInstanceStorageConfigCmd.MarkFlagRequired("storage-config")
	connectCmd.AddCommand(connect_updateInstanceStorageConfigCmd)
}
