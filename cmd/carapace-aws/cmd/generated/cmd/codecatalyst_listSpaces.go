package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_listSpacesCmd = &cobra.Command{
	Use:   "list-spaces",
	Short: "Retrieves a list of spaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_listSpacesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_listSpacesCmd).Standalone()

		codecatalyst_listSpacesCmd.Flags().String("next-token", "", "A token returned from a call to this API to indicate the next batch of results to return, if any.")
	})
	codecatalystCmd.AddCommand(codecatalyst_listSpacesCmd)
}
