package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createVerifiedAccessInstanceCmd = &cobra.Command{
	Use:   "create-verified-access-instance",
	Short: "An Amazon Web Services Verified Access instance is a regional entity that evaluates application requests and grants access only when your security requirements are met.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createVerifiedAccessInstanceCmd).Standalone()

	ec2_createVerifiedAccessInstanceCmd.Flags().String("cidr-endpoints-custom-sub-domain", "", "The custom subdomain.")
	ec2_createVerifiedAccessInstanceCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure idempotency of your modification request.")
	ec2_createVerifiedAccessInstanceCmd.Flags().String("description", "", "A description for the Verified Access instance.")
	ec2_createVerifiedAccessInstanceCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createVerifiedAccessInstanceCmd.Flags().Bool("fipsenabled", false, "Enable or disable support for Federal Information Processing Standards (FIPS) on the instance.")
	ec2_createVerifiedAccessInstanceCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createVerifiedAccessInstanceCmd.Flags().Bool("no-fipsenabled", false, "Enable or disable support for Federal Information Processing Standards (FIPS) on the instance.")
	ec2_createVerifiedAccessInstanceCmd.Flags().String("tag-specifications", "", "The tags to assign to the Verified Access instance.")
	ec2_createVerifiedAccessInstanceCmd.Flag("no-dry-run").Hidden = true
	ec2_createVerifiedAccessInstanceCmd.Flag("no-fipsenabled").Hidden = true
	ec2Cmd.AddCommand(ec2_createVerifiedAccessInstanceCmd)
}
