package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeHoursOfOperationOverrideCmd = &cobra.Command{
	Use:   "describe-hours-of-operation-override",
	Short: "Describes the hours of operation override.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeHoursOfOperationOverrideCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_describeHoursOfOperationOverrideCmd).Standalone()

		connect_describeHoursOfOperationOverrideCmd.Flags().String("hours-of-operation-id", "", "The identifier for the hours of operation.")
		connect_describeHoursOfOperationOverrideCmd.Flags().String("hours-of-operation-override-id", "", "The identifier for the hours of operation override.")
		connect_describeHoursOfOperationOverrideCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_describeHoursOfOperationOverrideCmd.MarkFlagRequired("hours-of-operation-id")
		connect_describeHoursOfOperationOverrideCmd.MarkFlagRequired("hours-of-operation-override-id")
		connect_describeHoursOfOperationOverrideCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_describeHoursOfOperationOverrideCmd)
}
