package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_createReferenceStoreCmd = &cobra.Command{
	Use:   "create-reference-store",
	Short: "Creates a reference store and returns metadata in JSON format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_createReferenceStoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_createReferenceStoreCmd).Standalone()

		omics_createReferenceStoreCmd.Flags().String("client-token", "", "To ensure that requests don't run multiple times, specify a unique token for each request.")
		omics_createReferenceStoreCmd.Flags().String("description", "", "A description for the store.")
		omics_createReferenceStoreCmd.Flags().String("name", "", "A name for the store.")
		omics_createReferenceStoreCmd.Flags().String("sse-config", "", "Server-side encryption (SSE) settings for the store.")
		omics_createReferenceStoreCmd.Flags().String("tags", "", "Tags for the store.")
		omics_createReferenceStoreCmd.MarkFlagRequired("name")
	})
	omicsCmd.AddCommand(omics_createReferenceStoreCmd)
}
