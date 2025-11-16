package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disassociateAddressCmd = &cobra.Command{
	Use:   "disassociate-address",
	Short: "Disassociates an Elastic IP address from the instance or network interface it's associated with.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disassociateAddressCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_disassociateAddressCmd).Standalone()

		ec2_disassociateAddressCmd.Flags().String("association-id", "", "The association ID.")
		ec2_disassociateAddressCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disassociateAddressCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disassociateAddressCmd.Flags().String("public-ip", "", "Deprecated.")
		ec2_disassociateAddressCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_disassociateAddressCmd)
}
