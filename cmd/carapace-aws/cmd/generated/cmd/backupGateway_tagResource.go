package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tag the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_tagResourceCmd).Standalone()

	backupGateway_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to tag.")
	backupGateway_tagResourceCmd.Flags().String("tags", "", "A list of tags to assign to the resource.")
	backupGateway_tagResourceCmd.MarkFlagRequired("resource-arn")
	backupGateway_tagResourceCmd.MarkFlagRequired("tags")
	backupGatewayCmd.AddCommand(backupGateway_tagResourceCmd)
}
