package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_describeScraperCmd = &cobra.Command{
	Use:   "describe-scraper",
	Short: "The `DescribeScraper` operation displays information about an existing scraper.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_describeScraperCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_describeScraperCmd).Standalone()

		amp_describeScraperCmd.Flags().String("scraper-id", "", "The ID of the scraper to describe.")
		amp_describeScraperCmd.MarkFlagRequired("scraper-id")
	})
	ampCmd.AddCommand(amp_describeScraperCmd)
}
