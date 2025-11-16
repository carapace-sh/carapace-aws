package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes a tag key and value from an AppConfig resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_untagResourceCmd).Standalone()

	appconfig_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource for which to remove tags.")
	appconfig_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys to delete.")
	appconfig_untagResourceCmd.MarkFlagRequired("resource-arn")
	appconfig_untagResourceCmd.MarkFlagRequired("tag-keys")
	appconfigCmd.AddCommand(appconfig_untagResourceCmd)
}
