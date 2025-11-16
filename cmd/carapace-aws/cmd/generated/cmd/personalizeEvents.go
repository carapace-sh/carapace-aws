package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalizeEventsCmd = &cobra.Command{
	Use:   "personalize-events",
	Short: "Amazon Personalize can consume real-time user event data, such as *stream* or *click* data, and use it for model training either alone or combined with historical data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalizeEventsCmd).Standalone()

	rootCmd.AddCommand(personalizeEventsCmd)
}
