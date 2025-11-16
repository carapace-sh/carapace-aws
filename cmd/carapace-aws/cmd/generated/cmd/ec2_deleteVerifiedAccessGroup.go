package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteVerifiedAccessGroupCmd = &cobra.Command{
	Use:   "delete-verified-access-group",
	Short: "Delete an Amazon Web Services Verified Access group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteVerifiedAccessGroupCmd).Standalone()

	ec2_deleteVerifiedAccessGroupCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure idempotency of your modification request.")
	ec2_deleteVerifiedAccessGroupCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteVerifiedAccessGroupCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteVerifiedAccessGroupCmd.Flags().String("verified-access-group-id", "", "The ID of the Verified Access group.")
	ec2_deleteVerifiedAccessGroupCmd.Flag("no-dry-run").Hidden = true
	ec2_deleteVerifiedAccessGroupCmd.MarkFlagRequired("verified-access-group-id")
	ec2Cmd.AddCommand(ec2_deleteVerifiedAccessGroupCmd)
}
