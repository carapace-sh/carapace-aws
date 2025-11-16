package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_deprecateActivityTypeCmd = &cobra.Command{
	Use:   "deprecate-activity-type",
	Short: "Deprecates the specified *activity type*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_deprecateActivityTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(swf_deprecateActivityTypeCmd).Standalone()

		swf_deprecateActivityTypeCmd.Flags().String("activity-type", "", "The activity type to deprecate.")
		swf_deprecateActivityTypeCmd.Flags().String("domain", "", "The name of the domain in which the activity type is registered.")
		swf_deprecateActivityTypeCmd.MarkFlagRequired("activity-type")
		swf_deprecateActivityTypeCmd.MarkFlagRequired("domain")
	})
	swfCmd.AddCommand(swf_deprecateActivityTypeCmd)
}
