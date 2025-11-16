package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listBotsCmd = &cobra.Command{
	Use:   "list-bots",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listBotsCmd).Standalone()

	connect_listBotsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_listBotsCmd.Flags().String("lex-version", "", "The version of Amazon Lex or Amazon Lex V2.")
	connect_listBotsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_listBotsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_listBotsCmd.MarkFlagRequired("instance-id")
	connect_listBotsCmd.MarkFlagRequired("lex-version")
	connectCmd.AddCommand(connect_listBotsCmd)
}
