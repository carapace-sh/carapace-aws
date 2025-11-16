package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3vectors_queryVectorsCmd = &cobra.Command{
	Use:   "query-vectors",
	Short: "Amazon S3 Vectors is in preview release for Amazon S3 and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3vectors_queryVectorsCmd).Standalone()

	s3vectors_queryVectorsCmd.Flags().String("filter", "", "Metadata filter to apply during the query.")
	s3vectors_queryVectorsCmd.Flags().String("index-arn", "", "The ARN of the vector index that you want to query.")
	s3vectors_queryVectorsCmd.Flags().String("index-name", "", "The name of the vector index that you want to query.")
	s3vectors_queryVectorsCmd.Flags().Bool("no-return-distance", false, "Indicates whether to include the computed distance in the response.")
	s3vectors_queryVectorsCmd.Flags().Bool("no-return-metadata", false, "Indicates whether to include metadata in the response.")
	s3vectors_queryVectorsCmd.Flags().String("query-vector", "", "The query vector.")
	s3vectors_queryVectorsCmd.Flags().Bool("return-distance", false, "Indicates whether to include the computed distance in the response.")
	s3vectors_queryVectorsCmd.Flags().Bool("return-metadata", false, "Indicates whether to include metadata in the response.")
	s3vectors_queryVectorsCmd.Flags().String("top-k", "", "The number of results to return for each query.")
	s3vectors_queryVectorsCmd.Flags().String("vector-bucket-name", "", "The name of the vector bucket that contains the vector index.")
	s3vectors_queryVectorsCmd.Flag("no-return-distance").Hidden = true
	s3vectors_queryVectorsCmd.Flag("no-return-metadata").Hidden = true
	s3vectors_queryVectorsCmd.MarkFlagRequired("query-vector")
	s3vectors_queryVectorsCmd.MarkFlagRequired("top-k")
	s3vectorsCmd.AddCommand(s3vectors_queryVectorsCmd)
}
