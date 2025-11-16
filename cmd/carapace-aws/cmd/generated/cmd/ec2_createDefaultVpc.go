package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createDefaultVpcCmd = &cobra.Command{
	Use:   "create-default-vpc",
	Short: "Creates a default VPC with a size `/16` IPv4 CIDR block and a default subnet in each Availability Zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createDefaultVpcCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createDefaultVpcCmd).Standalone()

		ec2_createDefaultVpcCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createDefaultVpcCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createDefaultVpcCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_createDefaultVpcCmd)
}
