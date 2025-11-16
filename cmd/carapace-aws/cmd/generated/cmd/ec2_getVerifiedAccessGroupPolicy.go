package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getVerifiedAccessGroupPolicyCmd = &cobra.Command{
	Use:   "get-verified-access-group-policy",
	Short: "Shows the contents of the Verified Access policy associated with the group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getVerifiedAccessGroupPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getVerifiedAccessGroupPolicyCmd).Standalone()

		ec2_getVerifiedAccessGroupPolicyCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getVerifiedAccessGroupPolicyCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getVerifiedAccessGroupPolicyCmd.Flags().String("verified-access-group-id", "", "The ID of the Verified Access group.")
		ec2_getVerifiedAccessGroupPolicyCmd.Flag("no-dry-run").Hidden = true
		ec2_getVerifiedAccessGroupPolicyCmd.MarkFlagRequired("verified-access-group-id")
	})
	ec2Cmd.AddCommand(ec2_getVerifiedAccessGroupPolicyCmd)
}
