package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_listObjectsCmd = &cobra.Command{
	Use:   "list-objects",
	Short: "End of support notice: Beginning November 21, 2025, Amazon S3 will stop returning `DisplayName`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_listObjectsCmd).Standalone()

	s3_listObjectsCmd.Flags().String("bucket", "", "The name of the bucket containing the objects.")
	s3_listObjectsCmd.Flags().String("delimiter", "", "A delimiter is a character that you use to group keys.")
	s3_listObjectsCmd.Flags().String("encoding-type", "", "")
	s3_listObjectsCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_listObjectsCmd.Flags().String("marker", "", "Marker is where you want Amazon S3 to start listing from.")
	s3_listObjectsCmd.Flags().String("max-keys", "", "Sets the maximum number of keys returned in the response.")
	s3_listObjectsCmd.Flags().String("optional-object-attributes", "", "Specifies the optional fields that you want returned in the response.")
	s3_listObjectsCmd.Flags().String("prefix", "", "Limits the response to keys that begin with the specified prefix.")
	s3_listObjectsCmd.Flags().String("request-payer", "", "Confirms that the requester knows that she or he will be charged for the list objects request.")
	s3_listObjectsCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_listObjectsCmd)
}
