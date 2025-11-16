package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getActiveNamesCmd = &cobra.Command{
	Use:   "get-active-names",
	Short: "Returns the names of all active (not deleted) resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getActiveNamesCmd).Standalone()

	lightsail_getActiveNamesCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	lightsailCmd.AddCommand(lightsail_getActiveNamesCmd)
}
