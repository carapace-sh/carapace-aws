package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2Cmd = &cobra.Command{
	Use:   "mediapackagev2",
	Short: "This guide is intended for creating AWS Elemental MediaPackage resources in MediaPackage Version 2 (v2) starting from May 2023. To get started with MediaPackage v2, create your MediaPackage resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackagev2Cmd).Standalone()

	})
	rootCmd.AddCommand(mediapackagev2Cmd)
}
