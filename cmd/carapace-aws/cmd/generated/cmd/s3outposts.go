package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3outpostsCmd = &cobra.Command{
	Use:   "s3outposts",
	Short: "Amazon S3 on Outposts provides access to S3 on Outposts operations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3outpostsCmd).Standalone()

	rootCmd.AddCommand(s3outpostsCmd)
}
