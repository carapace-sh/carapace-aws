package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_deleteRepositoryCmd = &cobra.Command{
	Use:   "delete-repository",
	Short: "De-register and unlink your repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_deleteRepositoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_deleteRepositoryCmd).Standalone()

		proton_deleteRepositoryCmd.Flags().String("name", "", "The repository name.")
		proton_deleteRepositoryCmd.Flags().String("provider", "", "The repository provider.")
		proton_deleteRepositoryCmd.MarkFlagRequired("name")
		proton_deleteRepositoryCmd.MarkFlagRequired("provider")
	})
	protonCmd.AddCommand(proton_deleteRepositoryCmd)
}
