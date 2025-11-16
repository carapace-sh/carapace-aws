package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_deleteScheduledActionCmd = &cobra.Command{
	Use:   "delete-scheduled-action",
	Short: "Deletes a scheduled action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_deleteScheduledActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_deleteScheduledActionCmd).Standalone()

		redshiftServerless_deleteScheduledActionCmd.Flags().String("scheduled-action-name", "", "The name of the scheduled action to delete.")
		redshiftServerless_deleteScheduledActionCmd.MarkFlagRequired("scheduled-action-name")
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_deleteScheduledActionCmd)
}
