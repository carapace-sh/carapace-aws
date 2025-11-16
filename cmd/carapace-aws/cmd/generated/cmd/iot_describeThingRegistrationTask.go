package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeThingRegistrationTaskCmd = &cobra.Command{
	Use:   "describe-thing-registration-task",
	Short: "Describes a bulk thing provisioning task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeThingRegistrationTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_describeThingRegistrationTaskCmd).Standalone()

		iot_describeThingRegistrationTaskCmd.Flags().String("task-id", "", "The task ID.")
		iot_describeThingRegistrationTaskCmd.MarkFlagRequired("task-id")
	})
	iotCmd.AddCommand(iot_describeThingRegistrationTaskCmd)
}
