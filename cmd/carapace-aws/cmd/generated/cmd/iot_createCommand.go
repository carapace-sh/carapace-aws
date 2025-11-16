package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createCommandCmd = &cobra.Command{
	Use:   "create-command",
	Short: "Creates a command.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createCommandCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_createCommandCmd).Standalone()

		iot_createCommandCmd.Flags().String("command-id", "", "A unique identifier for the command.")
		iot_createCommandCmd.Flags().String("description", "", "A short text decription of the command.")
		iot_createCommandCmd.Flags().String("display-name", "", "The user-friendly name in the console for the command.")
		iot_createCommandCmd.Flags().String("mandatory-parameters", "", "A list of parameters that are required by the `StartCommandExecution` API.")
		iot_createCommandCmd.Flags().String("namespace", "", "The namespace of the command.")
		iot_createCommandCmd.Flags().String("payload", "", "The payload object for the command.")
		iot_createCommandCmd.Flags().String("role-arn", "", "The IAM role that you must provide when using the `AWS-IoT-FleetWise` namespace.")
		iot_createCommandCmd.Flags().String("tags", "", "Name-value pairs that are used as metadata to manage a command.")
		iot_createCommandCmd.MarkFlagRequired("command-id")
	})
	iotCmd.AddCommand(iot_createCommandCmd)
}
