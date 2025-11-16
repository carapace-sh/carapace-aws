package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotData_getRetainedMessageCmd = &cobra.Command{
	Use:   "get-retained-message",
	Short: "Gets the details of a single retained message for the specified topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotData_getRetainedMessageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotData_getRetainedMessageCmd).Standalone()

		iotData_getRetainedMessageCmd.Flags().String("topic", "", "The topic name of the retained message to retrieve.")
		iotData_getRetainedMessageCmd.MarkFlagRequired("topic")
	})
	iotDataCmd.AddCommand(iotData_getRetainedMessageCmd)
}
