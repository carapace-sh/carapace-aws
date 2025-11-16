package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disassociateTrunkInterfaceCmd = &cobra.Command{
	Use:   "disassociate-trunk-interface",
	Short: "Removes an association between a branch network interface with a trunk network interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disassociateTrunkInterfaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_disassociateTrunkInterfaceCmd).Standalone()

		ec2_disassociateTrunkInterfaceCmd.Flags().String("association-id", "", "The ID of the association")
		ec2_disassociateTrunkInterfaceCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_disassociateTrunkInterfaceCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disassociateTrunkInterfaceCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disassociateTrunkInterfaceCmd.MarkFlagRequired("association-id")
		ec2_disassociateTrunkInterfaceCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_disassociateTrunkInterfaceCmd)
}
