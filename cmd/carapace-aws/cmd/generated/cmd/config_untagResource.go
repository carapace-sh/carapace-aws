package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes specified tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_untagResourceCmd).Standalone()

	config_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that identifies the resource for which to list the tags.")
	config_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags to be removed.")
	config_untagResourceCmd.MarkFlagRequired("resource-arn")
	config_untagResourceCmd.MarkFlagRequired("tag-keys")
	configCmd.AddCommand(config_untagResourceCmd)
}
