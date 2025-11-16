package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideoCmd = &cobra.Command{
	Use:   "kinesisvideo",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideoCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisvideoCmd).Standalone()

	})
	rootCmd.AddCommand(kinesisvideoCmd)
}
