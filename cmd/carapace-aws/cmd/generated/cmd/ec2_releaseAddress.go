package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_releaseAddressCmd = &cobra.Command{
	Use:   "release-address",
	Short: "Releases the specified Elastic IP address.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_releaseAddressCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_releaseAddressCmd).Standalone()

		ec2_releaseAddressCmd.Flags().String("allocation-id", "", "The allocation ID.")
		ec2_releaseAddressCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_releaseAddressCmd.Flags().String("network-border-group", "", "The set of Availability Zones, Local Zones, or Wavelength Zones from which Amazon Web Services advertises IP addresses.")
		ec2_releaseAddressCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_releaseAddressCmd.Flags().String("public-ip", "", "Deprecated.")
		ec2_releaseAddressCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_releaseAddressCmd)
}
