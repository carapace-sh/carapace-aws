package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listInstanceStorageConfigsCmd = &cobra.Command{
	Use:   "list-instance-storage-configs",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listInstanceStorageConfigsCmd).Standalone()

	connect_listInstanceStorageConfigsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_listInstanceStorageConfigsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_listInstanceStorageConfigsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_listInstanceStorageConfigsCmd.Flags().String("resource-type", "", "A valid resource type.")
	connect_listInstanceStorageConfigsCmd.MarkFlagRequired("instance-id")
	connect_listInstanceStorageConfigsCmd.MarkFlagRequired("resource-type")
	connectCmd.AddCommand(connect_listInstanceStorageConfigsCmd)
}
