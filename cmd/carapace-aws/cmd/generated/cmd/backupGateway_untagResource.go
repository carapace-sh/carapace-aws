package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_untagResourceCmd).Standalone()

	backupGateway_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource from which to remove tags.")
	backupGateway_untagResourceCmd.Flags().String("tag-keys", "", "The list of tag keys specifying which tags to remove.")
	backupGateway_untagResourceCmd.MarkFlagRequired("resource-arn")
	backupGateway_untagResourceCmd.MarkFlagRequired("tag-keys")
	backupGatewayCmd.AddCommand(backupGateway_untagResourceCmd)
}
