package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_getProtectedJobCmd = &cobra.Command{
	Use:   "get-protected-job",
	Short: "Returns job processing metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_getProtectedJobCmd).Standalone()

	cleanrooms_getProtectedJobCmd.Flags().String("membership-identifier", "", "The identifier for a membership in a protected job instance.")
	cleanrooms_getProtectedJobCmd.Flags().String("protected-job-identifier", "", "The identifier for the protected job instance.")
	cleanrooms_getProtectedJobCmd.MarkFlagRequired("membership-identifier")
	cleanrooms_getProtectedJobCmd.MarkFlagRequired("protected-job-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_getProtectedJobCmd)
}
