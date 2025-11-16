package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteHoursOfOperationOverrideCmd = &cobra.Command{
	Use:   "delete-hours-of-operation-override",
	Short: "Deletes an hours of operation override in an Amazon Connect hours of operation resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteHoursOfOperationOverrideCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_deleteHoursOfOperationOverrideCmd).Standalone()

		connect_deleteHoursOfOperationOverrideCmd.Flags().String("hours-of-operation-id", "", "The identifier for the hours of operation.")
		connect_deleteHoursOfOperationOverrideCmd.Flags().String("hours-of-operation-override-id", "", "The identifier for the hours of operation override.")
		connect_deleteHoursOfOperationOverrideCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_deleteHoursOfOperationOverrideCmd.MarkFlagRequired("hours-of-operation-id")
		connect_deleteHoursOfOperationOverrideCmd.MarkFlagRequired("hours-of-operation-override-id")
		connect_deleteHoursOfOperationOverrideCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_deleteHoursOfOperationOverrideCmd)
}
