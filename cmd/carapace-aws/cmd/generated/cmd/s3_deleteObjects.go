package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_deleteObjectsCmd = &cobra.Command{
	Use:   "delete-objects",
	Short: "This operation enables you to delete multiple objects from a bucket using a single HTTP request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_deleteObjectsCmd).Standalone()

	s3_deleteObjectsCmd.Flags().String("bucket", "", "The bucket name containing the objects to delete.")
	s3_deleteObjectsCmd.Flags().String("bypass-governance-retention", "", "Specifies whether you want to delete this object even if it has a Governance-type Object Lock in place.")
	s3_deleteObjectsCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the object when you use the SDK.")
	s3_deleteObjectsCmd.Flags().String("delete", "", "Container for the request.")
	s3_deleteObjectsCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_deleteObjectsCmd.Flags().String("mfa", "", "The concatenation of the authentication device's serial number, a space, and the value that is displayed on your authentication device.")
	s3_deleteObjectsCmd.Flags().String("request-payer", "", "")
	s3_deleteObjectsCmd.MarkFlagRequired("bucket")
	s3_deleteObjectsCmd.MarkFlagRequired("delete")
	s3Cmd.AddCommand(s3_deleteObjectsCmd)
}
