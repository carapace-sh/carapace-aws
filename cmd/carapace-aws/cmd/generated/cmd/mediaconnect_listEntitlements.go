package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_listEntitlementsCmd = &cobra.Command{
	Use:   "list-entitlements",
	Short: "Displays a list of all entitlements that have been granted to this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_listEntitlementsCmd).Standalone()

	mediaconnect_listEntitlementsCmd.Flags().String("max-results", "", "The maximum number of results to return per API request.")
	mediaconnect_listEntitlementsCmd.Flags().String("next-token", "", "The token that identifies the batch of results that you want to see.")
	mediaconnectCmd.AddCommand(mediaconnect_listEntitlementsCmd)
}
