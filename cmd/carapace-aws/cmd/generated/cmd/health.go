package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Health\n\nThe Health API provides access to the Health information that appears in the [Health Dashboard](https://health.aws.amazon.com/health/home). You can use the API operations to get information about events that might affect your Amazon Web Services services and resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(healthCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(healthCmd).Standalone()

	})
	rootCmd.AddCommand(healthCmd)
}
