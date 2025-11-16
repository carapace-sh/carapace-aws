package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyVerifiedAccessInstanceLoggingConfigurationCmd = &cobra.Command{
	Use:   "modify-verified-access-instance-logging-configuration",
	Short: "Modifies the logging configuration for the specified Amazon Web Services Verified Access instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyVerifiedAccessInstanceLoggingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyVerifiedAccessInstanceLoggingConfigurationCmd).Standalone()

		ec2_modifyVerifiedAccessInstanceLoggingConfigurationCmd.Flags().String("access-logs", "", "The configuration options for Verified Access instances.")
		ec2_modifyVerifiedAccessInstanceLoggingConfigurationCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure idempotency of your modification request.")
		ec2_modifyVerifiedAccessInstanceLoggingConfigurationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVerifiedAccessInstanceLoggingConfigurationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyVerifiedAccessInstanceLoggingConfigurationCmd.Flags().String("verified-access-instance-id", "", "The ID of the Verified Access instance.")
		ec2_modifyVerifiedAccessInstanceLoggingConfigurationCmd.MarkFlagRequired("access-logs")
		ec2_modifyVerifiedAccessInstanceLoggingConfigurationCmd.Flag("no-dry-run").Hidden = true
		ec2_modifyVerifiedAccessInstanceLoggingConfigurationCmd.MarkFlagRequired("verified-access-instance-id")
	})
	ec2Cmd.AddCommand(ec2_modifyVerifiedAccessInstanceLoggingConfigurationCmd)
}
