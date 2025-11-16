package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3vectors_listVectorsCmd = &cobra.Command{
	Use:   "list-vectors",
	Short: "Amazon S3 Vectors is in preview release for Amazon S3 and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3vectors_listVectorsCmd).Standalone()

	s3vectors_listVectorsCmd.Flags().String("index-arn", "", "The Amazon resource Name (ARN) of the vector index.")
	s3vectors_listVectorsCmd.Flags().String("index-name", "", "The name of the vector index.")
	s3vectors_listVectorsCmd.Flags().String("max-results", "", "The maximum number of vectors to return on a page.")
	s3vectors_listVectorsCmd.Flags().String("next-token", "", "Pagination token from a previous request.")
	s3vectors_listVectorsCmd.Flags().Bool("no-return-data", false, "If true, the vector data of each vector will be included in the response.")
	s3vectors_listVectorsCmd.Flags().Bool("no-return-metadata", false, "If true, the metadata associated with each vector will be included in the response.")
	s3vectors_listVectorsCmd.Flags().Bool("return-data", false, "If true, the vector data of each vector will be included in the response.")
	s3vectors_listVectorsCmd.Flags().Bool("return-metadata", false, "If true, the metadata associated with each vector will be included in the response.")
	s3vectors_listVectorsCmd.Flags().String("segment-count", "", "For a parallel `ListVectors` request, `segmentCount` represents the total number of vector segments into which the `ListVectors` operation will be divided.")
	s3vectors_listVectorsCmd.Flags().String("segment-index", "", "For a parallel `ListVectors` request, `segmentIndex` is the index of the segment from which to list vectors in the current request.")
	s3vectors_listVectorsCmd.Flags().String("vector-bucket-name", "", "The name of the vector bucket.")
	s3vectors_listVectorsCmd.Flag("no-return-data").Hidden = true
	s3vectors_listVectorsCmd.Flag("no-return-metadata").Hidden = true
	s3vectorsCmd.AddCommand(s3vectors_listVectorsCmd)
}
