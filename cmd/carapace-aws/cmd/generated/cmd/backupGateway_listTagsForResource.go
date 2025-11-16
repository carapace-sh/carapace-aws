package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags applied to the resource identified by its Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_listTagsForResourceCmd).Standalone()

	backupGateway_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource's tags to list.")
	backupGateway_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	backupGatewayCmd.AddCommand(backupGateway_listTagsForResourceCmd)
}
