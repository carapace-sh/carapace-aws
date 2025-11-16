package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_withdrawByoipCidrCmd = &cobra.Command{
	Use:   "withdraw-byoip-cidr",
	Short: "Stops advertising an address range that is provisioned as an address pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_withdrawByoipCidrCmd).Standalone()

	ec2_withdrawByoipCidrCmd.Flags().String("cidr", "", "The address range, in CIDR notation.")
	ec2_withdrawByoipCidrCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_withdrawByoipCidrCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_withdrawByoipCidrCmd.MarkFlagRequired("cidr")
	ec2_withdrawByoipCidrCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_withdrawByoipCidrCmd)
}
