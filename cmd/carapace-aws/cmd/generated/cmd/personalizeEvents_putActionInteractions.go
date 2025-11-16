package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalizeEvents_putActionInteractionsCmd = &cobra.Command{
	Use:   "put-action-interactions",
	Short: "Records action interaction event data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalizeEvents_putActionInteractionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalizeEvents_putActionInteractionsCmd).Standalone()

		personalizeEvents_putActionInteractionsCmd.Flags().String("action-interactions", "", "A list of action interaction events from the session.")
		personalizeEvents_putActionInteractionsCmd.Flags().String("tracking-id", "", "The ID of your action interaction event tracker.")
		personalizeEvents_putActionInteractionsCmd.MarkFlagRequired("action-interactions")
		personalizeEvents_putActionInteractionsCmd.MarkFlagRequired("tracking-id")
	})
	personalizeEventsCmd.AddCommand(personalizeEvents_putActionInteractionsCmd)
}
