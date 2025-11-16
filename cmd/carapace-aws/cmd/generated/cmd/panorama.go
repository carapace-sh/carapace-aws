package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panoramaCmd = &cobra.Command{
	Use:   "panorama",
	Short: "AWS Panorama\n\n**Overview**\n\nThis is the *AWS Panorama API Reference*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panoramaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(panoramaCmd).Standalone()

	})
	rootCmd.AddCommand(panoramaCmd)
}
