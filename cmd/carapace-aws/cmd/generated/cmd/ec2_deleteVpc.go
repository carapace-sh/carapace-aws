package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteVpcCmd = &cobra.Command{
	Use:   "delete-vpc",
	Short: "Deletes the specified VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteVpcCmd).Standalone()

	ec2_deleteVpcCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteVpcCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteVpcCmd.Flags().String("vpc-id", "", "The ID of the VPC.")
	ec2_deleteVpcCmd.Flag("no-dry-run").Hidden = true
	ec2_deleteVpcCmd.MarkFlagRequired("vpc-id")
	ec2Cmd.AddCommand(ec2_deleteVpcCmd)
}
