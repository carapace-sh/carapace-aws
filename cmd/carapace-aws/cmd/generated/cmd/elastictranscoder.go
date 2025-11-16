package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elastictranscoderCmd = &cobra.Command{
	Use:   "elastictranscoder",
	Short: "AWS Elastic Transcoder Service\n\nThe AWS Elastic Transcoder Service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elastictranscoderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elastictranscoderCmd).Standalone()

	})
	rootCmd.AddCommand(elastictranscoderCmd)
}
