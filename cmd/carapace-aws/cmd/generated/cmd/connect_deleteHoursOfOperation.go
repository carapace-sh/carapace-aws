package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteHoursOfOperationCmd = &cobra.Command{
	Use:   "delete-hours-of-operation",
	Short: "Deletes an hours of operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteHoursOfOperationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_deleteHoursOfOperationCmd).Standalone()

		connect_deleteHoursOfOperationCmd.Flags().String("hours-of-operation-id", "", "The identifier for the hours of operation.")
		connect_deleteHoursOfOperationCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_deleteHoursOfOperationCmd.MarkFlagRequired("hours-of-operation-id")
		connect_deleteHoursOfOperationCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_deleteHoursOfOperationCmd)
}
