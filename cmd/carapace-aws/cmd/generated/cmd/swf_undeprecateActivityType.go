package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_undeprecateActivityTypeCmd = &cobra.Command{
	Use:   "undeprecate-activity-type",
	Short: "Undeprecates a previously deprecated *activity type*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_undeprecateActivityTypeCmd).Standalone()

	swf_undeprecateActivityTypeCmd.Flags().String("activity-type", "", "The activity type to undeprecate.")
	swf_undeprecateActivityTypeCmd.Flags().String("domain", "", "The name of the domain of the deprecated activity type.")
	swf_undeprecateActivityTypeCmd.MarkFlagRequired("activity-type")
	swf_undeprecateActivityTypeCmd.MarkFlagRequired("domain")
	swfCmd.AddCommand(swf_undeprecateActivityTypeCmd)
}
