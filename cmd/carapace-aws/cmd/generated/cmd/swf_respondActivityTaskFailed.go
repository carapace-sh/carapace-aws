package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_respondActivityTaskFailedCmd = &cobra.Command{
	Use:   "respond-activity-task-failed",
	Short: "Used by workers to tell the service that the [ActivityTask]() identified by the `taskToken` has failed with `reason` (if specified).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_respondActivityTaskFailedCmd).Standalone()

	swf_respondActivityTaskFailedCmd.Flags().String("details", "", "Detailed information about the failure.")
	swf_respondActivityTaskFailedCmd.Flags().String("reason", "", "Description of the error that may assist in diagnostics.")
	swf_respondActivityTaskFailedCmd.Flags().String("task-token", "", "The `taskToken` of the [ActivityTask]().")
	swf_respondActivityTaskFailedCmd.MarkFlagRequired("task-token")
	swfCmd.AddCommand(swf_respondActivityTaskFailedCmd)
}
