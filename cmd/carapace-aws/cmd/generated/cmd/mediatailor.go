package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailorCmd = &cobra.Command{
	Use:   "mediatailor",
	Short: "Use the AWS Elemental MediaTailor SDKs and CLI to configure scalable ad insertion and linear channels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailorCmd).Standalone()

	})
	rootCmd.AddCommand(mediatailorCmd)
}
