package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_listObjectVersionsCmd = &cobra.Command{
	Use:   "list-object-versions",
	Short: "End of support notice: Beginning November 21, 2025, Amazon S3 will stop returning `DisplayName`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_listObjectVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_listObjectVersionsCmd).Standalone()

		s3_listObjectVersionsCmd.Flags().String("bucket", "", "The bucket name that contains the objects.")
		s3_listObjectVersionsCmd.Flags().String("delimiter", "", "A delimiter is a character that you specify to group keys.")
		s3_listObjectVersionsCmd.Flags().String("encoding-type", "", "")
		s3_listObjectVersionsCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_listObjectVersionsCmd.Flags().String("key-marker", "", "Specifies the key to start with when listing objects in a bucket.")
		s3_listObjectVersionsCmd.Flags().String("max-keys", "", "Sets the maximum number of keys returned in the response.")
		s3_listObjectVersionsCmd.Flags().String("optional-object-attributes", "", "Specifies the optional fields that you want returned in the response.")
		s3_listObjectVersionsCmd.Flags().String("prefix", "", "Use this parameter to select only those keys that begin with the specified prefix.")
		s3_listObjectVersionsCmd.Flags().String("request-payer", "", "")
		s3_listObjectVersionsCmd.Flags().String("version-id-marker", "", "Specifies the object version you want to start listing from.")
		s3_listObjectVersionsCmd.MarkFlagRequired("bucket")
	})
	s3Cmd.AddCommand(s3_listObjectVersionsCmd)
}
