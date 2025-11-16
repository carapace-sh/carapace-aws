package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_updateScraperCmd = &cobra.Command{
	Use:   "update-scraper",
	Short: "Updates an existing scraper.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_updateScraperCmd).Standalone()

	amp_updateScraperCmd.Flags().String("alias", "", "The new alias of the scraper.")
	amp_updateScraperCmd.Flags().String("client-token", "", "A unique identifier that you can provide to ensure the idempotency of the request.")
	amp_updateScraperCmd.Flags().String("destination", "", "The new Amazon Managed Service for Prometheus workspace to send metrics to.")
	amp_updateScraperCmd.Flags().String("role-configuration", "", "Use this structure to enable cross-account access, so that you can use a target account to access Prometheus metrics from source accounts.")
	amp_updateScraperCmd.Flags().String("scrape-configuration", "", "Contains the base-64 encoded YAML configuration for the scraper.")
	amp_updateScraperCmd.Flags().String("scraper-id", "", "The ID of the scraper to update.")
	amp_updateScraperCmd.MarkFlagRequired("scraper-id")
	ampCmd.AddCommand(amp_updateScraperCmd)
}
