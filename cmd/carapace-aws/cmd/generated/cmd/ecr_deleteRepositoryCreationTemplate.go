package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_deleteRepositoryCreationTemplateCmd = &cobra.Command{
	Use:   "delete-repository-creation-template",
	Short: "Deletes a repository creation template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_deleteRepositoryCreationTemplateCmd).Standalone()

	ecr_deleteRepositoryCreationTemplateCmd.Flags().String("prefix", "", "The repository namespace prefix associated with the repository creation template.")
	ecr_deleteRepositoryCreationTemplateCmd.MarkFlagRequired("prefix")
	ecrCmd.AddCommand(ecr_deleteRepositoryCreationTemplateCmd)
}
