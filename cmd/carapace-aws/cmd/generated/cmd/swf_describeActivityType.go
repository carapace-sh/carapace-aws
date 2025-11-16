package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_describeActivityTypeCmd = &cobra.Command{
	Use:   "describe-activity-type",
	Short: "Returns information about the specified activity type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_describeActivityTypeCmd).Standalone()

	swf_describeActivityTypeCmd.Flags().String("activity-type", "", "The activity type to get information about.")
	swf_describeActivityTypeCmd.Flags().String("domain", "", "The name of the domain in which the activity type is registered.")
	swf_describeActivityTypeCmd.MarkFlagRequired("activity-type")
	swf_describeActivityTypeCmd.MarkFlagRequired("domain")
	swfCmd.AddCommand(swf_describeActivityTypeCmd)
}
