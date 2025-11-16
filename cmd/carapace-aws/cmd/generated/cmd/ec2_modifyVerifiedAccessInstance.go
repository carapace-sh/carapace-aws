package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyVerifiedAccessInstanceCmd = &cobra.Command{
	Use:   "modify-verified-access-instance",
	Short: "Modifies the configuration of the specified Amazon Web Services Verified Access instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyVerifiedAccessInstanceCmd).Standalone()

	ec2_modifyVerifiedAccessInstanceCmd.Flags().String("cidr-endpoints-custom-sub-domain", "", "The custom subdomain.")
	ec2_modifyVerifiedAccessInstanceCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure idempotency of your modification request.")
	ec2_modifyVerifiedAccessInstanceCmd.Flags().String("description", "", "A description for the Verified Access instance.")
	ec2_modifyVerifiedAccessInstanceCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyVerifiedAccessInstanceCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyVerifiedAccessInstanceCmd.Flags().String("verified-access-instance-id", "", "The ID of the Verified Access instance.")
	ec2_modifyVerifiedAccessInstanceCmd.Flag("no-dry-run").Hidden = true
	ec2_modifyVerifiedAccessInstanceCmd.MarkFlagRequired("verified-access-instance-id")
	ec2Cmd.AddCommand(ec2_modifyVerifiedAccessInstanceCmd)
}
