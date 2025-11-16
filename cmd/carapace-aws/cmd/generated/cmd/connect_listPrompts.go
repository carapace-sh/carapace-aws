package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listPromptsCmd = &cobra.Command{
	Use:   "list-prompts",
	Short: "Provides information about the prompts for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listPromptsCmd).Standalone()

	connect_listPromptsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_listPromptsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_listPromptsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_listPromptsCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_listPromptsCmd)
}
