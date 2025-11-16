package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeAddressesAttributeCmd = &cobra.Command{
	Use:   "describe-addresses-attribute",
	Short: "Describes the attributes of the specified Elastic IP addresses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeAddressesAttributeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeAddressesAttributeCmd).Standalone()

		ec2_describeAddressesAttributeCmd.Flags().String("allocation-ids", "", "\\[EC2-VPC] The allocation IDs.")
		ec2_describeAddressesAttributeCmd.Flags().String("attribute", "", "The attribute of the IP address.")
		ec2_describeAddressesAttributeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeAddressesAttributeCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeAddressesAttributeCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeAddressesAttributeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeAddressesAttributeCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeAddressesAttributeCmd)
}
