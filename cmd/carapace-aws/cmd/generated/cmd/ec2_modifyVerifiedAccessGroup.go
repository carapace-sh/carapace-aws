package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyVerifiedAccessGroupCmd = &cobra.Command{
	Use:   "modify-verified-access-group",
	Short: "Modifies the specified Amazon Web Services Verified Access group configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyVerifiedAccessGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyVerifiedAccessGroupCmd).Standalone()

		ec2_modifyVerifiedAccessGroupCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure idempotency of your modification request.")
		ec2_modifyVerifiedAccessGroupCmd.Flags().String("description", "", "A description for the Verified Access group.")
		ec2_modifyVerifiedAccessGroupCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVerifiedAccessGroupCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVerifiedAccessGroupCmd.Flags().String("verified-access-group-id", "", "The ID of the Verified Access group.")
		ec2_modifyVerifiedAccessGroupCmd.Flags().String("verified-access-instance-id", "", "The ID of the Verified Access instance.")
		ec2_modifyVerifiedAccessGroupCmd.Flag("no-dry-run").Hidden = true
		ec2_modifyVerifiedAccessGroupCmd.MarkFlagRequired("verified-access-group-id")
	})
	ec2Cmd.AddCommand(ec2_modifyVerifiedAccessGroupCmd)
}
