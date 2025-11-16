package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendCmd = &cobra.Command{
	Use:   "comprehend",
	Short: "Amazon Comprehend is an Amazon Web Services service for gaining insight into the content of documents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehendCmd).Standalone()

	})
	rootCmd.AddCommand(comprehendCmd)
}
