package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_stopThingRegistrationTaskCmd = &cobra.Command{
	Use:   "stop-thing-registration-task",
	Short: "Cancels a bulk thing provisioning task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_stopThingRegistrationTaskCmd).Standalone()

	iot_stopThingRegistrationTaskCmd.Flags().String("task-id", "", "The bulk thing provisioning task ID.")
	iot_stopThingRegistrationTaskCmd.MarkFlagRequired("task-id")
	iotCmd.AddCommand(iot_stopThingRegistrationTaskCmd)
}
