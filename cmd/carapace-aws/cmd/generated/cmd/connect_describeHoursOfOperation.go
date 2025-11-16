package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeHoursOfOperationCmd = &cobra.Command{
	Use:   "describe-hours-of-operation",
	Short: "Describes the hours of operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeHoursOfOperationCmd).Standalone()

	connect_describeHoursOfOperationCmd.Flags().String("hours-of-operation-id", "", "The identifier for the hours of operation.")
	connect_describeHoursOfOperationCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_describeHoursOfOperationCmd.MarkFlagRequired("hours-of-operation-id")
	connect_describeHoursOfOperationCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_describeHoursOfOperationCmd)
}
