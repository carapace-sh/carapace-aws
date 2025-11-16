package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_putEmailMonitoringConfigurationCmd = &cobra.Command{
	Use:   "put-email-monitoring-configuration",
	Short: "Creates or updates the email monitoring configuration for a specified organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_putEmailMonitoringConfigurationCmd).Standalone()

	workmail_putEmailMonitoringConfigurationCmd.Flags().String("log-group-arn", "", "The Amazon Resource Name (ARN) of the CloudWatch Log group associated with the email monitoring configuration.")
	workmail_putEmailMonitoringConfigurationCmd.Flags().String("organization-id", "", "The ID of the organization for which the email monitoring configuration is set.")
	workmail_putEmailMonitoringConfigurationCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM Role associated with the email monitoring configuration.")
	workmail_putEmailMonitoringConfigurationCmd.MarkFlagRequired("log-group-arn")
	workmail_putEmailMonitoringConfigurationCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_putEmailMonitoringConfigurationCmd)
}
