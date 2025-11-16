package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3vectors_createIndexCmd = &cobra.Command{
	Use:   "create-index",
	Short: "Amazon S3 Vectors is in preview release for Amazon S3 and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3vectors_createIndexCmd).Standalone()

	s3vectors_createIndexCmd.Flags().String("data-type", "", "The data type of the vectors to be inserted into the vector index.")
	s3vectors_createIndexCmd.Flags().String("dimension", "", "The dimensions of the vectors to be inserted into the vector index.")
	s3vectors_createIndexCmd.Flags().String("distance-metric", "", "The distance metric to be used for similarity search.")
	s3vectors_createIndexCmd.Flags().String("index-name", "", "The name of the vector index to create.")
	s3vectors_createIndexCmd.Flags().String("metadata-configuration", "", "The metadata configuration for the vector index.")
	s3vectors_createIndexCmd.Flags().String("vector-bucket-arn", "", "The Amazon Resource Name (ARN) of the vector bucket to create the vector index in.")
	s3vectors_createIndexCmd.Flags().String("vector-bucket-name", "", "The name of the vector bucket to create the vector index in.")
	s3vectors_createIndexCmd.MarkFlagRequired("data-type")
	s3vectors_createIndexCmd.MarkFlagRequired("dimension")
	s3vectors_createIndexCmd.MarkFlagRequired("distance-metric")
	s3vectors_createIndexCmd.MarkFlagRequired("index-name")
	s3vectorsCmd.AddCommand(s3vectors_createIndexCmd)
}
