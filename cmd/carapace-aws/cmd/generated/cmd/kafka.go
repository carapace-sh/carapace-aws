package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "The operations for managing an Amazon MSK cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafkaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafkaCmd).Standalone()

	})
	rootCmd.AddCommand(kafkaCmd)
}
