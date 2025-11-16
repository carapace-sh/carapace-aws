package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_replaceRouteTableAssociationCmd = &cobra.Command{
	Use:   "replace-route-table-association",
	Short: "Changes the route table associated with a given subnet, internet gateway, or virtual private gateway in a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_replaceRouteTableAssociationCmd).Standalone()

	ec2_replaceRouteTableAssociationCmd.Flags().String("association-id", "", "The association ID.")
	ec2_replaceRouteTableAssociationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_replaceRouteTableAssociationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_replaceRouteTableAssociationCmd.Flags().String("route-table-id", "", "The ID of the new route table to associate with the subnet.")
	ec2_replaceRouteTableAssociationCmd.MarkFlagRequired("association-id")
	ec2_replaceRouteTableAssociationCmd.Flag("no-dry-run").Hidden = true
	ec2_replaceRouteTableAssociationCmd.MarkFlagRequired("route-table-id")
	ec2Cmd.AddCommand(ec2_replaceRouteTableAssociationCmd)
}
