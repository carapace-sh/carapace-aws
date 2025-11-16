package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_listLaunchActionsCmd = &cobra.Command{
	Use:   "list-launch-actions",
	Short: "Lists resource launch actions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_listLaunchActionsCmd).Standalone()

	drs_listLaunchActionsCmd.Flags().String("filters", "", "Filters to apply when listing resource launch actions.")
	drs_listLaunchActionsCmd.Flags().String("max-results", "", "Maximum amount of items to return when listing resource launch actions.")
	drs_listLaunchActionsCmd.Flags().String("next-token", "", "Next token to use when listing resource launch actions.")
	drs_listLaunchActionsCmd.Flags().String("resource-id", "", "")
	drs_listLaunchActionsCmd.MarkFlagRequired("resource-id")
	drsCmd.AddCommand(drs_listLaunchActionsCmd)
}
