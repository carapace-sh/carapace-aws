package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_deleteScraperLoggingConfigurationCmd = &cobra.Command{
	Use:   "delete-scraper-logging-configuration",
	Short: "Deletes the logging configuration for a Amazon Managed Service for Prometheus scraper.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_deleteScraperLoggingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_deleteScraperLoggingConfigurationCmd).Standalone()

		amp_deleteScraperLoggingConfigurationCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the request is processed exactly once.")
		amp_deleteScraperLoggingConfigurationCmd.Flags().String("scraper-id", "", "The ID of the scraper whose logging configuration will be deleted.")
		amp_deleteScraperLoggingConfigurationCmd.MarkFlagRequired("scraper-id")
	})
	ampCmd.AddCommand(amp_deleteScraperLoggingConfigurationCmd)
}
