package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createTransitGatewayCmd = &cobra.Command{
	Use:   "create-transit-gateway",
	Short: "Creates a transit gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createTransitGatewayCmd).Standalone()

	ec2_createTransitGatewayCmd.Flags().String("description", "", "A description of the transit gateway.")
	ec2_createTransitGatewayCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createTransitGatewayCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createTransitGatewayCmd.Flags().String("options", "", "The transit gateway options.")
	ec2_createTransitGatewayCmd.Flags().String("tag-specifications", "", "The tags to apply to the transit gateway.")
	ec2_createTransitGatewayCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_createTransitGatewayCmd)
}
