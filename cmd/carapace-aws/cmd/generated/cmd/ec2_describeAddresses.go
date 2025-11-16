package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeAddressesCmd = &cobra.Command{
	Use:   "describe-addresses",
	Short: "Describes the specified Elastic IP addresses or all of your Elastic IP addresses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeAddressesCmd).Standalone()

	ec2_describeAddressesCmd.Flags().String("allocation-ids", "", "Information about the allocation IDs.")
	ec2_describeAddressesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeAddressesCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeAddressesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeAddressesCmd.Flags().String("public-ips", "", "One or more Elastic IP addresses.")
	ec2_describeAddressesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeAddressesCmd)
}
