package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_deleteShareCmd = &cobra.Command{
	Use:   "delete-share",
	Short: "Deletes a resource share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_deleteShareCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_deleteShareCmd).Standalone()

		omics_deleteShareCmd.Flags().String("share-id", "", "The ID for the resource share to be deleted.")
		omics_deleteShareCmd.MarkFlagRequired("share-id")
	})
	omicsCmd.AddCommand(omics_deleteShareCmd)
}
