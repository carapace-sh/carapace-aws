package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_describeLagsCmd = &cobra.Command{
	Use:   "describe-lags",
	Short: "Describes all your link aggregation groups (LAG) or the specified LAG.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_describeLagsCmd).Standalone()

	directconnect_describeLagsCmd.Flags().String("lag-id", "", "The ID of the LAG.")
	directconnect_describeLagsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	directconnect_describeLagsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	directconnectCmd.AddCommand(directconnect_describeLagsCmd)
}
