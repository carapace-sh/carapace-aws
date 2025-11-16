package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_restoreObjectCmd = &cobra.Command{
	Use:   "restore-object",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_restoreObjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_restoreObjectCmd).Standalone()

		s3_restoreObjectCmd.Flags().String("bucket", "", "The bucket name containing the object to restore.")
		s3_restoreObjectCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the object when you use the SDK.")
		s3_restoreObjectCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_restoreObjectCmd.Flags().String("key", "", "Object key for which the action was initiated.")
		s3_restoreObjectCmd.Flags().String("request-payer", "", "")
		s3_restoreObjectCmd.Flags().String("restore-request", "", "")
		s3_restoreObjectCmd.Flags().String("version-id", "", "VersionId used to reference a specific version of the object.")
		s3_restoreObjectCmd.MarkFlagRequired("bucket")
		s3_restoreObjectCmd.MarkFlagRequired("key")
	})
	s3Cmd.AddCommand(s3_restoreObjectCmd)
}
