package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisVideoMediaCmd = &cobra.Command{
	Use:   "kinesis-video-media",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisVideoMediaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisVideoMediaCmd).Standalone()

	})
	rootCmd.AddCommand(kinesisVideoMediaCmd)
}
