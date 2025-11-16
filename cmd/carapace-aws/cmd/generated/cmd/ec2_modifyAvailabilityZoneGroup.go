package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyAvailabilityZoneGroupCmd = &cobra.Command{
	Use:   "modify-availability-zone-group",
	Short: "Changes the opt-in status of the specified zone group for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyAvailabilityZoneGroupCmd).Standalone()

	ec2_modifyAvailabilityZoneGroupCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyAvailabilityZoneGroupCmd.Flags().String("group-name", "", "The name of the Availability Zone group, Local Zone group, or Wavelength Zone group.")
	ec2_modifyAvailabilityZoneGroupCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyAvailabilityZoneGroupCmd.Flags().String("opt-in-status", "", "Indicates whether to opt in to the zone group.")
	ec2_modifyAvailabilityZoneGroupCmd.MarkFlagRequired("group-name")
	ec2_modifyAvailabilityZoneGroupCmd.Flag("no-dry-run").Hidden = true
	ec2_modifyAvailabilityZoneGroupCmd.MarkFlagRequired("opt-in-status")
	ec2Cmd.AddCommand(ec2_modifyAvailabilityZoneGroupCmd)
}
