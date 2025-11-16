package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3vectors_putVectorsCmd = &cobra.Command{
	Use:   "put-vectors",
	Short: "Amazon S3 Vectors is in preview release for Amazon S3 and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3vectors_putVectorsCmd).Standalone()

	s3vectors_putVectorsCmd.Flags().String("index-arn", "", "The ARN of the vector index where you want to write vectors.")
	s3vectors_putVectorsCmd.Flags().String("index-name", "", "The name of the vector index where you want to write vectors.")
	s3vectors_putVectorsCmd.Flags().String("vector-bucket-name", "", "The name of the vector bucket that contains the vector index.")
	s3vectors_putVectorsCmd.Flags().String("vectors", "", "The vectors to add to a vector index.")
	s3vectors_putVectorsCmd.MarkFlagRequired("vectors")
	s3vectorsCmd.AddCommand(s3vectors_putVectorsCmd)
}
