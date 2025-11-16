package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_describeEmailMonitoringConfigurationCmd = &cobra.Command{
	Use:   "describe-email-monitoring-configuration",
	Short: "Describes the current email monitoring configuration for a specified organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_describeEmailMonitoringConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_describeEmailMonitoringConfigurationCmd).Standalone()

		workmail_describeEmailMonitoringConfigurationCmd.Flags().String("organization-id", "", "The ID of the organization for which the email monitoring configuration is described.")
		workmail_describeEmailMonitoringConfigurationCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_describeEmailMonitoringConfigurationCmd)
}
