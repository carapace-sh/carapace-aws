package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createCoipCidrCmd = &cobra.Command{
	Use:   "create-coip-cidr",
	Short: "Creates a range of customer-owned IP addresses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createCoipCidrCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createCoipCidrCmd).Standalone()

		ec2_createCoipCidrCmd.Flags().String("cidr", "", "A customer-owned IP address range to create.")
		ec2_createCoipCidrCmd.Flags().String("coip-pool-id", "", "The ID of the address pool.")
		ec2_createCoipCidrCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createCoipCidrCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createCoipCidrCmd.MarkFlagRequired("cidr")
		ec2_createCoipCidrCmd.MarkFlagRequired("coip-pool-id")
		ec2_createCoipCidrCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_createCoipCidrCmd)
}
