package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyInstanceMaintenanceOptionsCmd = &cobra.Command{
	Use:   "modify-instance-maintenance-options",
	Short: "Modifies the recovery behavior of your instance to disable simplified automatic recovery or set the recovery behavior to default.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyInstanceMaintenanceOptionsCmd).Standalone()

	ec2_modifyInstanceMaintenanceOptionsCmd.Flags().String("auto-recovery", "", "Disables the automatic recovery behavior of your instance or sets it to default.")
	ec2_modifyInstanceMaintenanceOptionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyInstanceMaintenanceOptionsCmd.Flags().String("instance-id", "", "The ID of the instance.")
	ec2_modifyInstanceMaintenanceOptionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyInstanceMaintenanceOptionsCmd.Flags().String("reboot-migration", "", "Specifies whether to attempt reboot migration during a user-initiated reboot of an instance that has a scheduled `system-reboot` event:")
	ec2_modifyInstanceMaintenanceOptionsCmd.MarkFlagRequired("instance-id")
	ec2_modifyInstanceMaintenanceOptionsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_modifyInstanceMaintenanceOptionsCmd)
}
