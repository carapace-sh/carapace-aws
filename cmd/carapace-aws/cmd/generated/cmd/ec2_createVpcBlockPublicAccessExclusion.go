package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createVpcBlockPublicAccessExclusionCmd = &cobra.Command{
	Use:   "create-vpc-block-public-access-exclusion",
	Short: "Create a VPC Block Public Access (BPA) exclusion.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createVpcBlockPublicAccessExclusionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createVpcBlockPublicAccessExclusionCmd).Standalone()

		ec2_createVpcBlockPublicAccessExclusionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createVpcBlockPublicAccessExclusionCmd.Flags().String("internet-gateway-exclusion-mode", "", "The exclusion mode for internet gateway traffic.")
		ec2_createVpcBlockPublicAccessExclusionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createVpcBlockPublicAccessExclusionCmd.Flags().String("subnet-id", "", "A subnet ID.")
		ec2_createVpcBlockPublicAccessExclusionCmd.Flags().String("tag-specifications", "", "`tag` - The key/value combination of a tag assigned to the resource.")
		ec2_createVpcBlockPublicAccessExclusionCmd.Flags().String("vpc-id", "", "A VPC ID.")
		ec2_createVpcBlockPublicAccessExclusionCmd.MarkFlagRequired("internet-gateway-exclusion-mode")
		ec2_createVpcBlockPublicAccessExclusionCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_createVpcBlockPublicAccessExclusionCmd)
}
