package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listMultiplexesCmd = &cobra.Command{
	Use:   "list-multiplexes",
	Short: "Retrieve a list of the existing multiplexes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listMultiplexesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_listMultiplexesCmd).Standalone()

		medialive_listMultiplexesCmd.Flags().String("max-results", "", "The maximum number of items to return.")
		medialive_listMultiplexesCmd.Flags().String("next-token", "", "The token to retrieve the next page of results.")
	})
	medialiveCmd.AddCommand(medialive_listMultiplexesCmd)
}
