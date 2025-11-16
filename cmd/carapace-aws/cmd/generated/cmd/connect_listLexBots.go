package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listLexBotsCmd = &cobra.Command{
	Use:   "list-lex-bots",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listLexBotsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listLexBotsCmd).Standalone()

		connect_listLexBotsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listLexBotsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listLexBotsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listLexBotsCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_listLexBotsCmd)
}
