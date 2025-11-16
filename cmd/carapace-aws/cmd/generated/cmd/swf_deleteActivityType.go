package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_deleteActivityTypeCmd = &cobra.Command{
	Use:   "delete-activity-type",
	Short: "Deletes the specified *activity type*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_deleteActivityTypeCmd).Standalone()

	swf_deleteActivityTypeCmd.Flags().String("activity-type", "", "The activity type to delete.")
	swf_deleteActivityTypeCmd.Flags().String("domain", "", "The name of the domain in which the activity type is registered.")
	swf_deleteActivityTypeCmd.MarkFlagRequired("activity-type")
	swf_deleteActivityTypeCmd.MarkFlagRequired("domain")
	swfCmd.AddCommand(swf_deleteActivityTypeCmd)
}
