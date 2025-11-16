package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_putPublicAccessBlockCmd = &cobra.Command{
	Use:   "put-public-access-block",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_putPublicAccessBlockCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_putPublicAccessBlockCmd).Standalone()

		s3control_putPublicAccessBlockCmd.Flags().String("account-id", "", "The account ID for the Amazon Web Services account whose `PublicAccessBlock` configuration you want to set.")
		s3control_putPublicAccessBlockCmd.Flags().String("public-access-block-configuration", "", "The `PublicAccessBlock` configuration that you want to apply to the specified Amazon Web Services account.")
		s3control_putPublicAccessBlockCmd.MarkFlagRequired("account-id")
		s3control_putPublicAccessBlockCmd.MarkFlagRequired("public-access-block-configuration")
	})
	s3controlCmd.AddCommand(s3control_putPublicAccessBlockCmd)
}
