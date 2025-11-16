package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_moveAddressToVpcCmd = &cobra.Command{
	Use:   "move-address-to-vpc",
	Short: "This action is deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_moveAddressToVpcCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_moveAddressToVpcCmd).Standalone()

		ec2_moveAddressToVpcCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_moveAddressToVpcCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_moveAddressToVpcCmd.Flags().String("public-ip", "", "The Elastic IP address.")
		ec2_moveAddressToVpcCmd.Flag("no-dry-run").Hidden = true
		ec2_moveAddressToVpcCmd.MarkFlagRequired("public-ip")
	})
	ec2Cmd.AddCommand(ec2_moveAddressToVpcCmd)
}
