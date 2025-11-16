package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcRegionSwitch_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or updates tags for a Region switch resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcRegionSwitch_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(arcRegionSwitch_tagResourceCmd).Standalone()

		arcRegionSwitch_tagResourceCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) for a tag that you add to a resource.")
		arcRegionSwitch_tagResourceCmd.Flags().String("tags", "", "Tags that you add to a resource.")
		arcRegionSwitch_tagResourceCmd.MarkFlagRequired("arn")
		arcRegionSwitch_tagResourceCmd.MarkFlagRequired("tags")
	})
	arcRegionSwitchCmd.AddCommand(arcRegionSwitch_tagResourceCmd)
}
