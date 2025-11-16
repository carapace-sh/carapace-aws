package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafkaconnectCmd = &cobra.Command{
	Use:   "kafkaconnect",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafkaconnectCmd).Standalone()

	rootCmd.AddCommand(kafkaconnectCmd)
}
