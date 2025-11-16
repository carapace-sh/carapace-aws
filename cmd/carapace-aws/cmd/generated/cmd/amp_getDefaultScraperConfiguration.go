package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_getDefaultScraperConfigurationCmd = &cobra.Command{
	Use:   "get-default-scraper-configuration",
	Short: "The `GetDefaultScraperConfiguration` operation returns the default scraper configuration used when Amazon EKS creates a scraper for you.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_getDefaultScraperConfigurationCmd).Standalone()

	ampCmd.AddCommand(amp_getDefaultScraperConfigurationCmd)
}
