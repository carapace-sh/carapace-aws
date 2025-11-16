package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_getKxUserCmd = &cobra.Command{
	Use:   "get-kx-user",
	Short: "Retrieves information about the specified kdb user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_getKxUserCmd).Standalone()

	finspace_getKxUserCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment.")
	finspace_getKxUserCmd.Flags().String("user-name", "", "A unique identifier for the user.")
	finspace_getKxUserCmd.MarkFlagRequired("environment-id")
	finspace_getKxUserCmd.MarkFlagRequired("user-name")
	finspaceCmd.AddCommand(finspace_getKxUserCmd)
}
