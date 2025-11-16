package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_deleteAccessPointCmd = &cobra.Command{
	Use:   "delete-access-point",
	Short: "Deletes the specified access point.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_deleteAccessPointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(efs_deleteAccessPointCmd).Standalone()

		efs_deleteAccessPointCmd.Flags().String("access-point-id", "", "The ID of the access point that you want to delete.")
		efs_deleteAccessPointCmd.MarkFlagRequired("access-point-id")
	})
	efsCmd.AddCommand(efs_deleteAccessPointCmd)
}
