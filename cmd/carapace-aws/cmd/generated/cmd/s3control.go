package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3controlCmd = &cobra.Command{
	Use:   "s3control",
	Short: "Amazon Web Services S3 Control provides access to Amazon S3 control plane actions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3controlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3controlCmd).Standalone()

	})
	rootCmd.AddCommand(s3controlCmd)
}
