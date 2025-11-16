package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeMaintenanceWindowScheduleCmd = &cobra.Command{
	Use:   "describe-maintenance-window-schedule",
	Short: "Retrieves information about upcoming executions of a maintenance window.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeMaintenanceWindowScheduleCmd).Standalone()

	ssm_describeMaintenanceWindowScheduleCmd.Flags().String("filters", "", "Filters used to limit the range of results.")
	ssm_describeMaintenanceWindowScheduleCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_describeMaintenanceWindowScheduleCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	ssm_describeMaintenanceWindowScheduleCmd.Flags().String("resource-type", "", "The type of resource you want to retrieve information about.")
	ssm_describeMaintenanceWindowScheduleCmd.Flags().String("targets", "", "The managed node ID or key-value pair to retrieve information about.")
	ssm_describeMaintenanceWindowScheduleCmd.Flags().String("window-id", "", "The ID of the maintenance window to retrieve information about.")
	ssmCmd.AddCommand(ssm_describeMaintenanceWindowScheduleCmd)
}
