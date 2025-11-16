package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deprovisionByoipCidrCmd = &cobra.Command{
	Use:   "deprovision-byoip-cidr",
	Short: "Releases the specified address range that you provisioned for use with your Amazon Web Services resources through bring your own IP addresses (BYOIP) and deletes the corresponding address pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deprovisionByoipCidrCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deprovisionByoipCidrCmd).Standalone()

		ec2_deprovisionByoipCidrCmd.Flags().String("cidr", "", "The address range, in CIDR notation.")
		ec2_deprovisionByoipCidrCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deprovisionByoipCidrCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deprovisionByoipCidrCmd.MarkFlagRequired("cidr")
		ec2_deprovisionByoipCidrCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deprovisionByoipCidrCmd)
}
