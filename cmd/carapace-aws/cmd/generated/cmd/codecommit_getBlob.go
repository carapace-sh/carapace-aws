package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_getBlobCmd = &cobra.Command{
	Use:   "get-blob",
	Short: "Returns the base-64 encoded content of an individual blob in a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_getBlobCmd).Standalone()

	codecommit_getBlobCmd.Flags().String("blob-id", "", "The ID of the blob, which is its SHA-1 pointer.")
	codecommit_getBlobCmd.Flags().String("repository-name", "", "The name of the repository that contains the blob.")
	codecommit_getBlobCmd.MarkFlagRequired("blob-id")
	codecommit_getBlobCmd.MarkFlagRequired("repository-name")
	codecommitCmd.AddCommand(codecommit_getBlobCmd)
}
