package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_listAssistantsCmd = &cobra.Command{
	Use:   "list-assistants",
	Short: "Lists information about assistants.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_listAssistantsCmd).Standalone()

	wisdom_listAssistantsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	wisdom_listAssistantsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	wisdomCmd.AddCommand(wisdom_listAssistantsCmd)
}
