package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_getRepositoryTriggersCmd = &cobra.Command{
	Use:   "get-repository-triggers",
	Short: "Gets information about triggers configured for a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_getRepositoryTriggersCmd).Standalone()

	codecommit_getRepositoryTriggersCmd.Flags().String("repository-name", "", "The name of the repository for which the trigger is configured.")
	codecommit_getRepositoryTriggersCmd.MarkFlagRequired("repository-name")
	codecommitCmd.AddCommand(codecommit_getRepositoryTriggersCmd)
}
