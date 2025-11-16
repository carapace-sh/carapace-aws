package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_listRepositoriesCmd = &cobra.Command{
	Use:   "list-repositories",
	Short: "List linked repositories with detail data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_listRepositoriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_listRepositoriesCmd).Standalone()

		proton_listRepositoriesCmd.Flags().String("max-results", "", "The maximum number of repositories to list.")
		proton_listRepositoriesCmd.Flags().String("next-token", "", "A token that indicates the location of the next repository in the array of repositories, after the list of repositories previously requested.")
	})
	protonCmd.AddCommand(proton_listRepositoriesCmd)
}
