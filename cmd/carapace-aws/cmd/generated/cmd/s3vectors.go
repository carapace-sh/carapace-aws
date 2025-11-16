package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3vectorsCmd = &cobra.Command{
	Use:   "s3vectors",
	Short: "Amazon S3 vector buckets are a bucket type to store and search vectors with sub-second search times.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3vectorsCmd).Standalone()

	rootCmd.AddCommand(s3vectorsCmd)
}
