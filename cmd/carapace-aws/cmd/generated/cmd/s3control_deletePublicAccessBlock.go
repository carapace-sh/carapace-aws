package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_deletePublicAccessBlockCmd = &cobra.Command{
	Use:   "delete-public-access-block",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_deletePublicAccessBlockCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_deletePublicAccessBlockCmd).Standalone()

		s3control_deletePublicAccessBlockCmd.Flags().String("account-id", "", "The account ID for the Amazon Web Services account whose `PublicAccessBlock` configuration you want to remove.")
		s3control_deletePublicAccessBlockCmd.MarkFlagRequired("account-id")
	})
	s3controlCmd.AddCommand(s3control_deletePublicAccessBlockCmd)
}
