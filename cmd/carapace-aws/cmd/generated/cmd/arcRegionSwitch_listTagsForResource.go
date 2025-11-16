package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcRegionSwitch_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags attached to a Region switch resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcRegionSwitch_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(arcRegionSwitch_listTagsForResourceCmd).Standalone()

		arcRegionSwitch_listTagsForResourceCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the resource.")
		arcRegionSwitch_listTagsForResourceCmd.MarkFlagRequired("arn")
	})
	arcRegionSwitchCmd.AddCommand(arcRegionSwitch_listTagsForResourceCmd)
}
