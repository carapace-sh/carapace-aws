package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_getSignalMapCmd = &cobra.Command{
	Use:   "get-signal-map",
	Short: "Retrieves the specified signal map.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_getSignalMapCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_getSignalMapCmd).Standalone()

		medialive_getSignalMapCmd.Flags().String("identifier", "", "A signal map's identifier.")
		medialive_getSignalMapCmd.MarkFlagRequired("identifier")
	})
	medialiveCmd.AddCommand(medialive_getSignalMapCmd)
}
