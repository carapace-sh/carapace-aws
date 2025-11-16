package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeCmd = &cobra.Command{
	Use:   "chime",
	Short: "**Most of these APIs are no longer supported and will not be updated.** We recommend using the latest versions in the [Amazon Chime SDK API reference](https://docs.aws.amazon.com/chime-sdk/latest/APIReference/welcome.html), in the Amazon Chime SDK.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeCmd).Standalone()

	})
	rootCmd.AddCommand(chimeCmd)
}
