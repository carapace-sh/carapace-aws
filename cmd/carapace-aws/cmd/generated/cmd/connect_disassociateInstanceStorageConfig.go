package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_disassociateInstanceStorageConfigCmd = &cobra.Command{
	Use:   "disassociate-instance-storage-config",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_disassociateInstanceStorageConfigCmd).Standalone()

	connect_disassociateInstanceStorageConfigCmd.Flags().String("association-id", "", "The existing association identifier that uniquely identifies the resource type and storage config for the given instance ID.")
	connect_disassociateInstanceStorageConfigCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_disassociateInstanceStorageConfigCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_disassociateInstanceStorageConfigCmd.Flags().String("resource-type", "", "A valid resource type.")
	connect_disassociateInstanceStorageConfigCmd.MarkFlagRequired("association-id")
	connect_disassociateInstanceStorageConfigCmd.MarkFlagRequired("instance-id")
	connect_disassociateInstanceStorageConfigCmd.MarkFlagRequired("resource-type")
	connectCmd.AddCommand(connect_disassociateInstanceStorageConfigCmd)
}
