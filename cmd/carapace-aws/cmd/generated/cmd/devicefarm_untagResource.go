package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes the specified tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_untagResourceCmd).Standalone()

	devicefarm_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource or resources from which to delete tags.")
	devicefarm_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags to be removed.")
	devicefarm_untagResourceCmd.MarkFlagRequired("resource-arn")
	devicefarm_untagResourceCmd.MarkFlagRequired("tag-keys")
	devicefarmCmd.AddCommand(devicefarm_untagResourceCmd)
}
