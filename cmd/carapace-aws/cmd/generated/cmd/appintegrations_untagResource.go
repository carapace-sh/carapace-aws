package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrations_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrations_untagResourceCmd).Standalone()

	appintegrations_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	appintegrations_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys.")
	appintegrations_untagResourceCmd.MarkFlagRequired("resource-arn")
	appintegrations_untagResourceCmd.MarkFlagRequired("tag-keys")
	appintegrationsCmd.AddCommand(appintegrations_untagResourceCmd)
}
