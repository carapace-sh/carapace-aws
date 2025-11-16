package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteVpcEndpointsCmd = &cobra.Command{
	Use:   "delete-vpc-endpoints",
	Short: "Deletes the specified VPC endpoints.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteVpcEndpointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteVpcEndpointsCmd).Standalone()

		ec2_deleteVpcEndpointsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteVpcEndpointsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteVpcEndpointsCmd.Flags().String("vpc-endpoint-ids", "", "The IDs of the VPC endpoints.")
		ec2_deleteVpcEndpointsCmd.Flag("no-dry-run").Hidden = true
		ec2_deleteVpcEndpointsCmd.MarkFlagRequired("vpc-endpoint-ids")
	})
	ec2Cmd.AddCommand(ec2_deleteVpcEndpointsCmd)
}
