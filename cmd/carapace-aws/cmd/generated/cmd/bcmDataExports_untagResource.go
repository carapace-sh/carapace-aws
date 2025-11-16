package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmDataExports_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes tags associated with an existing data export definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmDataExports_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmDataExports_untagResourceCmd).Standalone()

		bcmDataExports_untagResourceCmd.Flags().String("resource-arn", "", "The unique identifier for the resource.")
		bcmDataExports_untagResourceCmd.Flags().String("resource-tag-keys", "", "The tag keys that are associated with the resource ARN.")
		bcmDataExports_untagResourceCmd.MarkFlagRequired("resource-arn")
		bcmDataExports_untagResourceCmd.MarkFlagRequired("resource-tag-keys")
	})
	bcmDataExportsCmd.AddCommand(bcmDataExports_untagResourceCmd)
}
