package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getObjectLegalHoldCmd = &cobra.Command{
	Use:   "get-object-legal-hold",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getObjectLegalHoldCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_getObjectLegalHoldCmd).Standalone()

		s3_getObjectLegalHoldCmd.Flags().String("bucket", "", "The bucket name containing the object whose legal hold status you want to retrieve.")
		s3_getObjectLegalHoldCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_getObjectLegalHoldCmd.Flags().String("key", "", "The key name for the object whose legal hold status you want to retrieve.")
		s3_getObjectLegalHoldCmd.Flags().String("request-payer", "", "")
		s3_getObjectLegalHoldCmd.Flags().String("version-id", "", "The version ID of the object whose legal hold status you want to retrieve.")
		s3_getObjectLegalHoldCmd.MarkFlagRequired("bucket")
		s3_getObjectLegalHoldCmd.MarkFlagRequired("key")
	})
	s3Cmd.AddCommand(s3_getObjectLegalHoldCmd)
}
