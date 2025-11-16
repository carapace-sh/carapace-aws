package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_getScheduledActionCmd = &cobra.Command{
	Use:   "get-scheduled-action",
	Short: "Returns information about a scheduled action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_getScheduledActionCmd).Standalone()

	redshiftServerless_getScheduledActionCmd.Flags().String("scheduled-action-name", "", "The name of the scheduled action.")
	redshiftServerless_getScheduledActionCmd.MarkFlagRequired("scheduled-action-name")
	redshiftServerlessCmd.AddCommand(redshiftServerless_getScheduledActionCmd)
}
