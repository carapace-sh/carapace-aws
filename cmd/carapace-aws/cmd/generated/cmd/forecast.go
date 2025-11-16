package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Provides APIs for creating and managing Amazon Forecast resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecastCmd).Standalone()

	rootCmd.AddCommand(forecastCmd)
}
