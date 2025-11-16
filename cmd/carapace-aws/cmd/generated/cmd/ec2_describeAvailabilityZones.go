package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeAvailabilityZonesCmd = &cobra.Command{
	Use:   "describe-availability-zones",
	Short: "Describes the Availability Zones, Local Zones, and Wavelength Zones that are available to you.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeAvailabilityZonesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeAvailabilityZonesCmd).Standalone()

		ec2_describeAvailabilityZonesCmd.Flags().Bool("all-availability-zones", false, "Include all Availability Zones, Local Zones, and Wavelength Zones regardless of your opt-in status.")
		ec2_describeAvailabilityZonesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeAvailabilityZonesCmd.Flags().String("filters", "", "The filters.")
		ec2_describeAvailabilityZonesCmd.Flags().Bool("no-all-availability-zones", false, "Include all Availability Zones, Local Zones, and Wavelength Zones regardless of your opt-in status.")
		ec2_describeAvailabilityZonesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeAvailabilityZonesCmd.Flags().String("zone-ids", "", "The IDs of the Availability Zones, Local Zones, and Wavelength Zones.")
		ec2_describeAvailabilityZonesCmd.Flags().String("zone-names", "", "The names of the Availability Zones, Local Zones, and Wavelength Zones.")
		ec2_describeAvailabilityZonesCmd.Flag("no-all-availability-zones").Hidden = true
		ec2_describeAvailabilityZonesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeAvailabilityZonesCmd)
}
