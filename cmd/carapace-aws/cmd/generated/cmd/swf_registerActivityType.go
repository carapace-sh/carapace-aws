package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_registerActivityTypeCmd = &cobra.Command{
	Use:   "register-activity-type",
	Short: "Registers a new *activity type* along with its configuration settings in the specified domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_registerActivityTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(swf_registerActivityTypeCmd).Standalone()

		swf_registerActivityTypeCmd.Flags().String("default-task-heartbeat-timeout", "", "If set, specifies the default maximum time before which a worker processing a task of this type must report progress by calling [RecordActivityTaskHeartbeat]().")
		swf_registerActivityTypeCmd.Flags().String("default-task-list", "", "If set, specifies the default task list to use for scheduling tasks of this activity type.")
		swf_registerActivityTypeCmd.Flags().String("default-task-priority", "", "The default task priority to assign to the activity type.")
		swf_registerActivityTypeCmd.Flags().String("default-task-schedule-to-close-timeout", "", "If set, specifies the default maximum duration for a task of this activity type.")
		swf_registerActivityTypeCmd.Flags().String("default-task-schedule-to-start-timeout", "", "If set, specifies the default maximum duration that a task of this activity type can wait before being assigned to a worker.")
		swf_registerActivityTypeCmd.Flags().String("default-task-start-to-close-timeout", "", "If set, specifies the default maximum duration that a worker can take to process tasks of this activity type.")
		swf_registerActivityTypeCmd.Flags().String("description", "", "A textual description of the activity type.")
		swf_registerActivityTypeCmd.Flags().String("domain", "", "The name of the domain in which this activity is to be registered.")
		swf_registerActivityTypeCmd.Flags().String("name", "", "The name of the activity type within the domain.")
		swf_registerActivityTypeCmd.Flags().String("version", "", "The version of the activity type.")
		swf_registerActivityTypeCmd.MarkFlagRequired("domain")
		swf_registerActivityTypeCmd.MarkFlagRequired("name")
		swf_registerActivityTypeCmd.MarkFlagRequired("version")
	})
	swfCmd.AddCommand(swf_registerActivityTypeCmd)
}
