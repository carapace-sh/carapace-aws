package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createVerifiedAccessGroupCmd = &cobra.Command{
	Use:   "create-verified-access-group",
	Short: "An Amazon Web Services Verified Access group is a collection of Amazon Web Services Verified Access endpoints who's associated applications have similar security requirements.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createVerifiedAccessGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createVerifiedAccessGroupCmd).Standalone()

		ec2_createVerifiedAccessGroupCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure idempotency of your modification request.")
		ec2_createVerifiedAccessGroupCmd.Flags().String("description", "", "A description for the Verified Access group.")
		ec2_createVerifiedAccessGroupCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createVerifiedAccessGroupCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createVerifiedAccessGroupCmd.Flags().String("policy-document", "", "The Verified Access policy document.")
		ec2_createVerifiedAccessGroupCmd.Flags().String("sse-specification", "", "The options for server side encryption.")
		ec2_createVerifiedAccessGroupCmd.Flags().String("tag-specifications", "", "The tags to assign to the Verified Access group.")
		ec2_createVerifiedAccessGroupCmd.Flags().String("verified-access-instance-id", "", "The ID of the Verified Access instance.")
		ec2_createVerifiedAccessGroupCmd.Flag("no-dry-run").Hidden = true
		ec2_createVerifiedAccessGroupCmd.MarkFlagRequired("verified-access-instance-id")
	})
	ec2Cmd.AddCommand(ec2_createVerifiedAccessGroupCmd)
}
