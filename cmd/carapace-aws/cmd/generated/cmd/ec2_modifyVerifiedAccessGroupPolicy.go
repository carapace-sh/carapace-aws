package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyVerifiedAccessGroupPolicyCmd = &cobra.Command{
	Use:   "modify-verified-access-group-policy",
	Short: "Modifies the specified Amazon Web Services Verified Access group policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyVerifiedAccessGroupPolicyCmd).Standalone()

	ec2_modifyVerifiedAccessGroupPolicyCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure idempotency of your modification request.")
	ec2_modifyVerifiedAccessGroupPolicyCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyVerifiedAccessGroupPolicyCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyVerifiedAccessGroupPolicyCmd.Flags().Bool("no-policy-enabled", false, "The status of the Verified Access policy.")
	ec2_modifyVerifiedAccessGroupPolicyCmd.Flags().String("policy-document", "", "The Verified Access policy document.")
	ec2_modifyVerifiedAccessGroupPolicyCmd.Flags().Bool("policy-enabled", false, "The status of the Verified Access policy.")
	ec2_modifyVerifiedAccessGroupPolicyCmd.Flags().String("sse-specification", "", "The options for server side encryption.")
	ec2_modifyVerifiedAccessGroupPolicyCmd.Flags().String("verified-access-group-id", "", "The ID of the Verified Access group.")
	ec2_modifyVerifiedAccessGroupPolicyCmd.Flag("no-dry-run").Hidden = true
	ec2_modifyVerifiedAccessGroupPolicyCmd.Flag("no-policy-enabled").Hidden = true
	ec2_modifyVerifiedAccessGroupPolicyCmd.MarkFlagRequired("verified-access-group-id")
	ec2Cmd.AddCommand(ec2_modifyVerifiedAccessGroupPolicyCmd)
}
