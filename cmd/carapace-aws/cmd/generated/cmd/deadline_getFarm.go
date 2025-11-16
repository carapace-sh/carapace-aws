package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_getFarmCmd = &cobra.Command{
	Use:   "get-farm",
	Short: "Get a farm.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_getFarmCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_getFarmCmd).Standalone()

		deadline_getFarmCmd.Flags().String("farm-id", "", "The farm ID of the farm.")
		deadline_getFarmCmd.MarkFlagRequired("farm-id")
	})
	deadlineCmd.AddCommand(deadline_getFarmCmd)
}
