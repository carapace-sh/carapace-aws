package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_putRepositoryTriggersCmd = &cobra.Command{
	Use:   "put-repository-triggers",
	Short: "Replaces all triggers for a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_putRepositoryTriggersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_putRepositoryTriggersCmd).Standalone()

		codecommit_putRepositoryTriggersCmd.Flags().String("repository-name", "", "The name of the repository where you want to create or update the trigger.")
		codecommit_putRepositoryTriggersCmd.Flags().String("triggers", "", "The JSON block of configuration information for each trigger.")
		codecommit_putRepositoryTriggersCmd.MarkFlagRequired("repository-name")
		codecommit_putRepositoryTriggersCmd.MarkFlagRequired("triggers")
	})
	codecommitCmd.AddCommand(codecommit_putRepositoryTriggersCmd)
}
