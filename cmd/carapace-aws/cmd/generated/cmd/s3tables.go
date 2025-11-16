package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tablesCmd = &cobra.Command{
	Use:   "s3tables",
	Short: "An Amazon S3 table represents a structured dataset consisting of tabular data in [Apache Parquet](https://parquet.apache.org/docs/) format and related metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tablesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3tablesCmd).Standalone()

	})
	rootCmd.AddCommand(s3tablesCmd)
}
