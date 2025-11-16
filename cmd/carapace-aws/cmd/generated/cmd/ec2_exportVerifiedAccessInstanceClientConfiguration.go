package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_exportVerifiedAccessInstanceClientConfigurationCmd = &cobra.Command{
	Use:   "export-verified-access-instance-client-configuration",
	Short: "Exports the client configuration for a Verified Access instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_exportVerifiedAccessInstanceClientConfigurationCmd).Standalone()

	ec2_exportVerifiedAccessInstanceClientConfigurationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_exportVerifiedAccessInstanceClientConfigurationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_exportVerifiedAccessInstanceClientConfigurationCmd.Flags().String("verified-access-instance-id", "", "The ID of the Verified Access instance.")
	ec2_exportVerifiedAccessInstanceClientConfigurationCmd.Flag("no-dry-run").Hidden = true
	ec2_exportVerifiedAccessInstanceClientConfigurationCmd.MarkFlagRequired("verified-access-instance-id")
	ec2Cmd.AddCommand(ec2_exportVerifiedAccessInstanceClientConfigurationCmd)
}
