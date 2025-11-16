package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_getListsMetadataCmd = &cobra.Command{
	Use:   "get-lists-metadata",
	Short: "Gets the metadata of either all the lists under the account or the specified list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_getListsMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_getListsMetadataCmd).Standalone()

		frauddetector_getListsMetadataCmd.Flags().String("max-results", "", "The maximum number of objects to return for the request.")
		frauddetector_getListsMetadataCmd.Flags().String("name", "", "The name of the list.")
		frauddetector_getListsMetadataCmd.Flags().String("next-token", "", "The next token for the subsequent request.")
	})
	frauddetectorCmd.AddCommand(frauddetector_getListsMetadataCmd)
}
