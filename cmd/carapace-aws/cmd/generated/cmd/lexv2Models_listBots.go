package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listBotsCmd = &cobra.Command{
	Use:   "list-bots",
	Short: "Gets a list of available bots.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listBotsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_listBotsCmd).Standalone()

		lexv2Models_listBotsCmd.Flags().String("filters", "", "Provides the specification of a filter used to limit the bots in the response to only those that match the filter specification.")
		lexv2Models_listBotsCmd.Flags().String("max-results", "", "The maximum number of bots to return in each page of results.")
		lexv2Models_listBotsCmd.Flags().String("next-token", "", "If the response from the `ListBots` operation contains more results than specified in the `maxResults` parameter, a token is returned in the response.")
		lexv2Models_listBotsCmd.Flags().String("sort-by", "", "Specifies sorting parameters for the list of bots.")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_listBotsCmd)
}
