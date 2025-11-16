package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnectCmd = &cobra.Command{
	Use:   "qconnect",
	Short: "- [Amazon Q actions](https://docs.aws.amazon.com/connect/latest/APIReference/API_Operations_Amazon_Q_Connect.html)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnectCmd).Standalone()

	})
	rootCmd.AddCommand(qconnectCmd)
}
