package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_updateProtectedJobCmd = &cobra.Command{
	Use:   "update-protected-job",
	Short: "Updates the processing of a currently running job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_updateProtectedJobCmd).Standalone()

	cleanrooms_updateProtectedJobCmd.Flags().String("membership-identifier", "", "The identifier for a member of a protected job instance.")
	cleanrooms_updateProtectedJobCmd.Flags().String("protected-job-identifier", "", "The identifier of the protected job to update.")
	cleanrooms_updateProtectedJobCmd.Flags().String("target-status", "", "The target status of a protected job.")
	cleanrooms_updateProtectedJobCmd.MarkFlagRequired("membership-identifier")
	cleanrooms_updateProtectedJobCmd.MarkFlagRequired("protected-job-identifier")
	cleanrooms_updateProtectedJobCmd.MarkFlagRequired("target-status")
	cleanroomsCmd.AddCommand(cleanrooms_updateProtectedJobCmd)
}
