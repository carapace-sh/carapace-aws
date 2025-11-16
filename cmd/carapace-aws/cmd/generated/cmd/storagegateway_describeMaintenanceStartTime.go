package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_describeMaintenanceStartTimeCmd = &cobra.Command{
	Use:   "describe-maintenance-start-time",
	Short: "Returns your gateway's maintenance window schedule information, with values for monthly or weekly cadence, specific day and time to begin maintenance, and which types of updates to apply.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_describeMaintenanceStartTimeCmd).Standalone()

	storagegateway_describeMaintenanceStartTimeCmd.Flags().String("gateway-arn", "", "")
	storagegateway_describeMaintenanceStartTimeCmd.MarkFlagRequired("gateway-arn")
	storagegatewayCmd.AddCommand(storagegateway_describeMaintenanceStartTimeCmd)
}
