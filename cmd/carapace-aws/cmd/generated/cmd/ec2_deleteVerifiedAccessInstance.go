package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteVerifiedAccessInstanceCmd = &cobra.Command{
	Use:   "delete-verified-access-instance",
	Short: "Delete an Amazon Web Services Verified Access instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteVerifiedAccessInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteVerifiedAccessInstanceCmd).Standalone()

		ec2_deleteVerifiedAccessInstanceCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure idempotency of your modification request.")
		ec2_deleteVerifiedAccessInstanceCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteVerifiedAccessInstanceCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteVerifiedAccessInstanceCmd.Flags().String("verified-access-instance-id", "", "The ID of the Verified Access instance.")
		ec2_deleteVerifiedAccessInstanceCmd.Flag("no-dry-run").Hidden = true
		ec2_deleteVerifiedAccessInstanceCmd.MarkFlagRequired("verified-access-instance-id")
	})
	ec2Cmd.AddCommand(ec2_deleteVerifiedAccessInstanceCmd)
}
