package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_deleteSdiSourceCmd = &cobra.Command{
	Use:   "delete-sdi-source",
	Short: "Delete an SdiSource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_deleteSdiSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_deleteSdiSourceCmd).Standalone()

		medialive_deleteSdiSourceCmd.Flags().String("sdi-source-id", "", "The ID of the SdiSource.")
		medialive_deleteSdiSourceCmd.MarkFlagRequired("sdi-source-id")
	})
	medialiveCmd.AddCommand(medialive_deleteSdiSourceCmd)
}
