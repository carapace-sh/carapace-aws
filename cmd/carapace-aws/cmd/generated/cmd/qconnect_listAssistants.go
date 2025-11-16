package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_listAssistantsCmd = &cobra.Command{
	Use:   "list-assistants",
	Short: "Lists information about assistants.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_listAssistantsCmd).Standalone()

	qconnect_listAssistantsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	qconnect_listAssistantsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	qconnectCmd.AddCommand(qconnect_listAssistantsCmd)
}
