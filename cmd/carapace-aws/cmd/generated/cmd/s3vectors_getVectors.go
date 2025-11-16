package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3vectors_getVectorsCmd = &cobra.Command{
	Use:   "get-vectors",
	Short: "Amazon S3 Vectors is in preview release for Amazon S3 and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3vectors_getVectorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3vectors_getVectorsCmd).Standalone()

		s3vectors_getVectorsCmd.Flags().String("index-arn", "", "The ARN of the vector index.")
		s3vectors_getVectorsCmd.Flags().String("index-name", "", "The name of the vector index.")
		s3vectors_getVectorsCmd.Flags().String("keys", "", "The names of the vectors you want to return attributes for.")
		s3vectors_getVectorsCmd.Flags().Bool("no-return-data", false, "Indicates whether to include the vector data in the response.")
		s3vectors_getVectorsCmd.Flags().Bool("no-return-metadata", false, "Indicates whether to include metadata in the response.")
		s3vectors_getVectorsCmd.Flags().Bool("return-data", false, "Indicates whether to include the vector data in the response.")
		s3vectors_getVectorsCmd.Flags().Bool("return-metadata", false, "Indicates whether to include metadata in the response.")
		s3vectors_getVectorsCmd.Flags().String("vector-bucket-name", "", "The name of the vector bucket that contains the vector index.")
		s3vectors_getVectorsCmd.MarkFlagRequired("keys")
		s3vectors_getVectorsCmd.Flag("no-return-data").Hidden = true
		s3vectors_getVectorsCmd.Flag("no-return-metadata").Hidden = true
	})
	s3vectorsCmd.AddCommand(s3vectors_getVectorsCmd)
}
