package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var braketCmd = &cobra.Command{
	Use:   "braket",
	Short: "The Amazon Braket API Reference provides information about the operations and structures supported by Amazon Braket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(braketCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(braketCmd).Standalone()

	})
	rootCmd.AddCommand(braketCmd)
}
