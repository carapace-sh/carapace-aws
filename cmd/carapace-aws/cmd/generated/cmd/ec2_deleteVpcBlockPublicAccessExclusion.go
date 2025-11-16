package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteVpcBlockPublicAccessExclusionCmd = &cobra.Command{
	Use:   "delete-vpc-block-public-access-exclusion",
	Short: "Delete a VPC Block Public Access (BPA) exclusion.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteVpcBlockPublicAccessExclusionCmd).Standalone()

	ec2_deleteVpcBlockPublicAccessExclusionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteVpcBlockPublicAccessExclusionCmd.Flags().String("exclusion-id", "", "The ID of the exclusion.")
	ec2_deleteVpcBlockPublicAccessExclusionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteVpcBlockPublicAccessExclusionCmd.MarkFlagRequired("exclusion-id")
	ec2_deleteVpcBlockPublicAccessExclusionCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_deleteVpcBlockPublicAccessExclusionCmd)
}
