package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteCoipCidrCmd = &cobra.Command{
	Use:   "delete-coip-cidr",
	Short: "Deletes a range of customer-owned IP addresses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteCoipCidrCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteCoipCidrCmd).Standalone()

		ec2_deleteCoipCidrCmd.Flags().String("cidr", "", "A customer-owned IP address range that you want to delete.")
		ec2_deleteCoipCidrCmd.Flags().String("coip-pool-id", "", "The ID of the customer-owned address pool.")
		ec2_deleteCoipCidrCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteCoipCidrCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteCoipCidrCmd.MarkFlagRequired("cidr")
		ec2_deleteCoipCidrCmd.MarkFlagRequired("coip-pool-id")
		ec2_deleteCoipCidrCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteCoipCidrCmd)
}
