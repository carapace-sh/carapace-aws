package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_deleteScraperCmd = &cobra.Command{
	Use:   "delete-scraper",
	Short: "The `DeleteScraper` operation deletes one scraper, and stops any metrics collection that the scraper performs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_deleteScraperCmd).Standalone()

	amp_deleteScraperCmd.Flags().String("client-token", "", "(Optional) A unique, case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	amp_deleteScraperCmd.Flags().String("scraper-id", "", "The ID of the scraper to delete.")
	amp_deleteScraperCmd.MarkFlagRequired("scraper-id")
	ampCmd.AddCommand(amp_deleteScraperCmd)
}
