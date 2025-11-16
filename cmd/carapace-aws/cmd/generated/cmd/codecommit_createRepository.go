package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_createRepositoryCmd = &cobra.Command{
	Use:   "create-repository",
	Short: "Creates a new, empty repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_createRepositoryCmd).Standalone()

	codecommit_createRepositoryCmd.Flags().String("kms-key-id", "", "The ID of the encryption key.")
	codecommit_createRepositoryCmd.Flags().String("repository-description", "", "A comment or description about the new repository.")
	codecommit_createRepositoryCmd.Flags().String("repository-name", "", "The name of the new repository to be created.")
	codecommit_createRepositoryCmd.Flags().String("tags", "", "One or more tag key-value pairs to use when tagging this repository.")
	codecommit_createRepositoryCmd.MarkFlagRequired("repository-name")
	codecommitCmd.AddCommand(codecommit_createRepositoryCmd)
}
