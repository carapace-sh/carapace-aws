package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_associateInstanceStorageConfigCmd = &cobra.Command{
	Use:   "associate-instance-storage-config",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_associateInstanceStorageConfigCmd).Standalone()

	connect_associateInstanceStorageConfigCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_associateInstanceStorageConfigCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_associateInstanceStorageConfigCmd.Flags().String("resource-type", "", "A valid resource type.")
	connect_associateInstanceStorageConfigCmd.Flags().String("storage-config", "", "A valid storage type.")
	connect_associateInstanceStorageConfigCmd.MarkFlagRequired("instance-id")
	connect_associateInstanceStorageConfigCmd.MarkFlagRequired("resource-type")
	connect_associateInstanceStorageConfigCmd.MarkFlagRequired("storage-config")
	connectCmd.AddCommand(connect_associateInstanceStorageConfigCmd)
}
