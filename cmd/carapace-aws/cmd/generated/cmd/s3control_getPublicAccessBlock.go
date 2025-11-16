package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getPublicAccessBlockCmd = &cobra.Command{
	Use:   "get-public-access-block",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getPublicAccessBlockCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_getPublicAccessBlockCmd).Standalone()

		s3control_getPublicAccessBlockCmd.Flags().String("account-id", "", "The account ID for the Amazon Web Services account whose `PublicAccessBlock` configuration you want to retrieve.")
		s3control_getPublicAccessBlockCmd.MarkFlagRequired("account-id")
	})
	s3controlCmd.AddCommand(s3control_getPublicAccessBlockCmd)
}
