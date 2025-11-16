package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_deleteLaunchActionCmd = &cobra.Command{
	Use:   "delete-launch-action",
	Short: "Deletes a resource launch action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_deleteLaunchActionCmd).Standalone()

	drs_deleteLaunchActionCmd.Flags().String("action-id", "", "")
	drs_deleteLaunchActionCmd.Flags().String("resource-id", "", "")
	drs_deleteLaunchActionCmd.MarkFlagRequired("action-id")
	drs_deleteLaunchActionCmd.MarkFlagRequired("resource-id")
	drsCmd.AddCommand(drs_deleteLaunchActionCmd)
}
