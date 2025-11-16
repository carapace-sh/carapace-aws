package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_listTopicsCmd = &cobra.Command{
	Use:   "list-topics",
	Short: "Returns a list of the requester's topics.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_listTopicsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sns_listTopicsCmd).Standalone()

		sns_listTopicsCmd.Flags().String("next-token", "", "Token returned by the previous `ListTopics` request.")
	})
	snsCmd.AddCommand(sns_listTopicsCmd)
}
