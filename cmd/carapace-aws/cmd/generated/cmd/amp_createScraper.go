package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_createScraperCmd = &cobra.Command{
	Use:   "create-scraper",
	Short: "The `CreateScraper` operation creates a scraper to collect metrics.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_createScraperCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_createScraperCmd).Standalone()

		amp_createScraperCmd.Flags().String("alias", "", "(optional) An alias to associate with the scraper.")
		amp_createScraperCmd.Flags().String("client-token", "", "(Optional) A unique, case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		amp_createScraperCmd.Flags().String("destination", "", "The Amazon Managed Service for Prometheus workspace to send metrics to.")
		amp_createScraperCmd.Flags().String("role-configuration", "", "Use this structure to enable cross-account access, so that you can use a target account to access Prometheus metrics from source accounts.")
		amp_createScraperCmd.Flags().String("scrape-configuration", "", "The configuration file to use in the new scraper.")
		amp_createScraperCmd.Flags().String("source", "", "The Amazon EKS or Amazon Web Services cluster from which the scraper will collect metrics.")
		amp_createScraperCmd.Flags().String("tags", "", "(Optional) The list of tag keys and values to associate with the scraper.")
		amp_createScraperCmd.MarkFlagRequired("destination")
		amp_createScraperCmd.MarkFlagRequired("scrape-configuration")
		amp_createScraperCmd.MarkFlagRequired("source")
	})
	ampCmd.AddCommand(amp_createScraperCmd)
}
