package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_getRepositoryCmd = &cobra.Command{
	Use:   "get-repository",
	Short: "Get detail data for a linked repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_getRepositoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_getRepositoryCmd).Standalone()

		proton_getRepositoryCmd.Flags().String("name", "", "The repository name, for example `myrepos/myrepo`.")
		proton_getRepositoryCmd.Flags().String("provider", "", "The repository provider.")
		proton_getRepositoryCmd.MarkFlagRequired("name")
		proton_getRepositoryCmd.MarkFlagRequired("provider")
	})
	protonCmd.AddCommand(proton_getRepositoryCmd)
}
