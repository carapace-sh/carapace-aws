package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyVpcBlockPublicAccessExclusionCmd = &cobra.Command{
	Use:   "modify-vpc-block-public-access-exclusion",
	Short: "Modify VPC Block Public Access (BPA) exclusions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyVpcBlockPublicAccessExclusionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyVpcBlockPublicAccessExclusionCmd).Standalone()

		ec2_modifyVpcBlockPublicAccessExclusionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVpcBlockPublicAccessExclusionCmd.Flags().String("exclusion-id", "", "The ID of an exclusion.")
		ec2_modifyVpcBlockPublicAccessExclusionCmd.Flags().String("internet-gateway-exclusion-mode", "", "The exclusion mode for internet gateway traffic.")
		ec2_modifyVpcBlockPublicAccessExclusionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVpcBlockPublicAccessExclusionCmd.MarkFlagRequired("exclusion-id")
		ec2_modifyVpcBlockPublicAccessExclusionCmd.MarkFlagRequired("internet-gateway-exclusion-mode")
		ec2_modifyVpcBlockPublicAccessExclusionCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_modifyVpcBlockPublicAccessExclusionCmd)
}
