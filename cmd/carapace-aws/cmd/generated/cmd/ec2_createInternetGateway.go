package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createInternetGatewayCmd = &cobra.Command{
	Use:   "create-internet-gateway",
	Short: "Creates an internet gateway for use with a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createInternetGatewayCmd).Standalone()

	ec2_createInternetGatewayCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createInternetGatewayCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createInternetGatewayCmd.Flags().String("tag-specifications", "", "The tags to assign to the internet gateway.")
	ec2_createInternetGatewayCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_createInternetGatewayCmd)
}
