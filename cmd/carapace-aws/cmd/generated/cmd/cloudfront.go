package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfrontCmd = &cobra.Command{
	Use:   "cloudfront",
	Short: "Amazon CloudFront",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfrontCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfrontCmd).Standalone()

	})
	rootCmd.AddCommand(cloudfrontCmd)
}
