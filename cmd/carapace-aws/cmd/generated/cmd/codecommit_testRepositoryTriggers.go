package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_testRepositoryTriggersCmd = &cobra.Command{
	Use:   "test-repository-triggers",
	Short: "Tests the functionality of repository triggers by sending information to the trigger target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_testRepositoryTriggersCmd).Standalone()

	codecommit_testRepositoryTriggersCmd.Flags().String("repository-name", "", "The name of the repository in which to test the triggers.")
	codecommit_testRepositoryTriggersCmd.Flags().String("triggers", "", "The list of triggers to test.")
	codecommit_testRepositoryTriggersCmd.MarkFlagRequired("repository-name")
	codecommit_testRepositoryTriggersCmd.MarkFlagRequired("triggers")
	codecommitCmd.AddCommand(codecommit_testRepositoryTriggersCmd)
}
