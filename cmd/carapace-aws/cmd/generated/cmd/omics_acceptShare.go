package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_acceptShareCmd = &cobra.Command{
	Use:   "accept-share",
	Short: "Accept a resource share request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_acceptShareCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_acceptShareCmd).Standalone()

		omics_acceptShareCmd.Flags().String("share-id", "", "The ID of the resource share.")
		omics_acceptShareCmd.MarkFlagRequired("share-id")
	})
	omicsCmd.AddCommand(omics_acceptShareCmd)
}
