package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_listBridgesCmd = &cobra.Command{
	Use:   "list-bridges",
	Short: "Displays a list of bridges that are associated with this account and an optionally specified Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_listBridgesCmd).Standalone()

	mediaconnect_listBridgesCmd.Flags().String("filter-arn", "", "Filter the list results to display only the bridges associated with the selected ARN.")
	mediaconnect_listBridgesCmd.Flags().String("max-results", "", "The maximum number of results to return per API request.")
	mediaconnect_listBridgesCmd.Flags().String("next-token", "", "The token that identifies the batch of results that you want to see.")
	mediaconnectCmd.AddCommand(mediaconnect_listBridgesCmd)
}
