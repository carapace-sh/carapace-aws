package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_deleteSignalMapCmd = &cobra.Command{
	Use:   "delete-signal-map",
	Short: "Deletes the specified signal map.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_deleteSignalMapCmd).Standalone()

	medialive_deleteSignalMapCmd.Flags().String("identifier", "", "A signal map's identifier.")
	medialive_deleteSignalMapCmd.MarkFlagRequired("identifier")
	medialiveCmd.AddCommand(medialive_deleteSignalMapCmd)
}
