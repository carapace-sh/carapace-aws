package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_getShareCmd = &cobra.Command{
	Use:   "get-share",
	Short: "Retrieves the metadata for the specified resource share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_getShareCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_getShareCmd).Standalone()

		omics_getShareCmd.Flags().String("share-id", "", "The ID of the share.")
		omics_getShareCmd.MarkFlagRequired("share-id")
	})
	omicsCmd.AddCommand(omics_getShareCmd)
}
