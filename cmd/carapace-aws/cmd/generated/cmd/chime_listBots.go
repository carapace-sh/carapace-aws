package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_listBotsCmd = &cobra.Command{
	Use:   "list-bots",
	Short: "Lists the bots associated with the administrator's Amazon Chime Enterprise account ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_listBotsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_listBotsCmd).Standalone()

		chime_listBotsCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
		chime_listBotsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		chime_listBotsCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
		chime_listBotsCmd.MarkFlagRequired("account-id")
	})
	chimeCmd.AddCommand(chime_listBotsCmd)
}
