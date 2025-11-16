package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_getLimitCmd = &cobra.Command{
	Use:   "get-limit",
	Short: "Gets information about a specific limit.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_getLimitCmd).Standalone()

	deadline_getLimitCmd.Flags().String("farm-id", "", "The unique identifier of the farm that contains the limit.")
	deadline_getLimitCmd.Flags().String("limit-id", "", "The unique identifier of the limit to return.")
	deadline_getLimitCmd.MarkFlagRequired("farm-id")
	deadline_getLimitCmd.MarkFlagRequired("limit-id")
	deadlineCmd.AddCommand(deadline_getLimitCmd)
}
