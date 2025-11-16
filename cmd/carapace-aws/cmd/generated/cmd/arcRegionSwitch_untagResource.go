package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcRegionSwitch_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from a Region switch resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcRegionSwitch_untagResourceCmd).Standalone()

	arcRegionSwitch_untagResourceCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) for a tag you remove a resource from.")
	arcRegionSwitch_untagResourceCmd.Flags().String("resource-tag-keys", "", "Tag keys that you remove from a resource.")
	arcRegionSwitch_untagResourceCmd.MarkFlagRequired("arn")
	arcRegionSwitch_untagResourceCmd.MarkFlagRequired("resource-tag-keys")
	arcRegionSwitchCmd.AddCommand(arcRegionSwitch_untagResourceCmd)
}
