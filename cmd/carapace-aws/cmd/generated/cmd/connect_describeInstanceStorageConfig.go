package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeInstanceStorageConfigCmd = &cobra.Command{
	Use:   "describe-instance-storage-config",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeInstanceStorageConfigCmd).Standalone()

	connect_describeInstanceStorageConfigCmd.Flags().String("association-id", "", "The existing association identifier that uniquely identifies the resource type and storage config for the given instance ID.")
	connect_describeInstanceStorageConfigCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_describeInstanceStorageConfigCmd.Flags().String("resource-type", "", "A valid resource type.")
	connect_describeInstanceStorageConfigCmd.MarkFlagRequired("association-id")
	connect_describeInstanceStorageConfigCmd.MarkFlagRequired("instance-id")
	connect_describeInstanceStorageConfigCmd.MarkFlagRequired("resource-type")
	connectCmd.AddCommand(connect_describeInstanceStorageConfigCmd)
}
