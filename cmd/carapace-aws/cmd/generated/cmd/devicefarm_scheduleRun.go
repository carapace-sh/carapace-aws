package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_scheduleRunCmd = &cobra.Command{
	Use:   "schedule-run",
	Short: "Schedules a run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_scheduleRunCmd).Standalone()

	devicefarm_scheduleRunCmd.Flags().String("app-arn", "", "The ARN of an application package to run tests against, created with [CreateUpload]().")
	devicefarm_scheduleRunCmd.Flags().String("configuration", "", "Information about the settings for the run to be scheduled.")
	devicefarm_scheduleRunCmd.Flags().String("device-pool-arn", "", "The ARN of the device pool for the run to be scheduled.")
	devicefarm_scheduleRunCmd.Flags().String("device-selection-configuration", "", "The filter criteria used to dynamically select a set of devices for a test run and the maximum number of devices to be included in the run.")
	devicefarm_scheduleRunCmd.Flags().String("execution-configuration", "", "Specifies configuration information about a test run, such as the execution timeout (in minutes).")
	devicefarm_scheduleRunCmd.Flags().String("name", "", "The name for the run to be scheduled.")
	devicefarm_scheduleRunCmd.Flags().String("project-arn", "", "The ARN of the project for the run to be scheduled.")
	devicefarm_scheduleRunCmd.Flags().String("test", "", "Information about the test for the run to be scheduled.")
	devicefarm_scheduleRunCmd.MarkFlagRequired("project-arn")
	devicefarm_scheduleRunCmd.MarkFlagRequired("test")
	devicefarmCmd.AddCommand(devicefarm_scheduleRunCmd)
}
