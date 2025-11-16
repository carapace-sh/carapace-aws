package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrailCmd = &cobra.Command{
	Use:   "cloudtrail",
	Short: "CloudTrail",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrailCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrailCmd).Standalone()

	})
	rootCmd.AddCommand(cloudtrailCmd)
}
