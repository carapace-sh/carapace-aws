package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createDefaultSubnetCmd = &cobra.Command{
	Use:   "create-default-subnet",
	Short: "Creates a default subnet with a size `/20` IPv4 CIDR block in the specified Availability Zone in your default VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createDefaultSubnetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createDefaultSubnetCmd).Standalone()

		ec2_createDefaultSubnetCmd.Flags().String("availability-zone", "", "The Availability Zone in which to create the default subnet.")
		ec2_createDefaultSubnetCmd.Flags().String("availability-zone-id", "", "The ID of the Availability Zone.")
		ec2_createDefaultSubnetCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createDefaultSubnetCmd.Flags().Bool("ipv6-native", false, "Indicates whether to create an IPv6 only subnet.")
		ec2_createDefaultSubnetCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createDefaultSubnetCmd.Flags().Bool("no-ipv6-native", false, "Indicates whether to create an IPv6 only subnet.")
		ec2_createDefaultSubnetCmd.Flag("no-dry-run").Hidden = true
		ec2_createDefaultSubnetCmd.Flag("no-ipv6-native").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_createDefaultSubnetCmd)
}
