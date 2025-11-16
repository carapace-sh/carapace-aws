package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_deleteScheduledActionCmd = &cobra.Command{
	Use:   "delete-scheduled-action",
	Short: "Deletes a scheduled action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_deleteScheduledActionCmd).Standalone()

	redshift_deleteScheduledActionCmd.Flags().String("scheduled-action-name", "", "The name of the scheduled action to delete.")
	redshift_deleteScheduledActionCmd.MarkFlagRequired("scheduled-action-name")
	redshiftCmd.AddCommand(redshift_deleteScheduledActionCmd)
}
