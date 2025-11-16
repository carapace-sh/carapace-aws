package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_updateMaintenanceWindowTargetCmd = &cobra.Command{
	Use:   "update-maintenance-window-target",
	Short: "Modifies the target of an existing maintenance window.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_updateMaintenanceWindowTargetCmd).Standalone()

	ssm_updateMaintenanceWindowTargetCmd.Flags().String("description", "", "An optional description for the update.")
	ssm_updateMaintenanceWindowTargetCmd.Flags().String("name", "", "A name for the update.")
	ssm_updateMaintenanceWindowTargetCmd.Flags().Bool("no-replace", false, "If `True`, then all fields that are required by the [RegisterTargetWithMaintenanceWindow]() operation are also required for this API request.")
	ssm_updateMaintenanceWindowTargetCmd.Flags().String("owner-information", "", "User-provided value that will be included in any Amazon CloudWatch Events events raised while running tasks for these targets in this maintenance window.")
	ssm_updateMaintenanceWindowTargetCmd.Flags().Bool("replace", false, "If `True`, then all fields that are required by the [RegisterTargetWithMaintenanceWindow]() operation are also required for this API request.")
	ssm_updateMaintenanceWindowTargetCmd.Flags().String("targets", "", "The targets to add or replace.")
	ssm_updateMaintenanceWindowTargetCmd.Flags().String("window-id", "", "The maintenance window ID with which to modify the target.")
	ssm_updateMaintenanceWindowTargetCmd.Flags().String("window-target-id", "", "The target ID to modify.")
	ssm_updateMaintenanceWindowTargetCmd.Flag("no-replace").Hidden = true
	ssm_updateMaintenanceWindowTargetCmd.MarkFlagRequired("window-id")
	ssm_updateMaintenanceWindowTargetCmd.MarkFlagRequired("window-target-id")
	ssmCmd.AddCommand(ssm_updateMaintenanceWindowTargetCmd)
}
