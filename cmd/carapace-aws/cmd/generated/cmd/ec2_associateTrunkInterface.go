package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_associateTrunkInterfaceCmd = &cobra.Command{
	Use:   "associate-trunk-interface",
	Short: "Associates a branch network interface with a trunk network interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_associateTrunkInterfaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_associateTrunkInterfaceCmd).Standalone()

		ec2_associateTrunkInterfaceCmd.Flags().String("branch-interface-id", "", "The ID of the branch network interface.")
		ec2_associateTrunkInterfaceCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_associateTrunkInterfaceCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_associateTrunkInterfaceCmd.Flags().String("gre-key", "", "The application key.")
		ec2_associateTrunkInterfaceCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_associateTrunkInterfaceCmd.Flags().String("trunk-interface-id", "", "The ID of the trunk network interface.")
		ec2_associateTrunkInterfaceCmd.Flags().String("vlan-id", "", "The ID of the VLAN.")
		ec2_associateTrunkInterfaceCmd.MarkFlagRequired("branch-interface-id")
		ec2_associateTrunkInterfaceCmd.Flag("no-dry-run").Hidden = true
		ec2_associateTrunkInterfaceCmd.MarkFlagRequired("trunk-interface-id")
	})
	ec2Cmd.AddCommand(ec2_associateTrunkInterfaceCmd)
}
