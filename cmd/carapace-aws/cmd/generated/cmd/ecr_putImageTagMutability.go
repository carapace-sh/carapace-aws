package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_putImageTagMutabilityCmd = &cobra.Command{
	Use:   "put-image-tag-mutability",
	Short: "Updates the image tag mutability settings for the specified repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_putImageTagMutabilityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_putImageTagMutabilityCmd).Standalone()

		ecr_putImageTagMutabilityCmd.Flags().String("image-tag-mutability", "", "The tag mutability setting for the repository.")
		ecr_putImageTagMutabilityCmd.Flags().String("image-tag-mutability-exclusion-filters", "", "Creates or updates a repository with filters that define which image tags can override the default image tag mutability setting.")
		ecr_putImageTagMutabilityCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry that contains the repository in which to update the image tag mutability settings.")
		ecr_putImageTagMutabilityCmd.Flags().String("repository-name", "", "The name of the repository in which to update the image tag mutability settings.")
		ecr_putImageTagMutabilityCmd.MarkFlagRequired("image-tag-mutability")
		ecr_putImageTagMutabilityCmd.MarkFlagRequired("repository-name")
	})
	ecrCmd.AddCommand(ecr_putImageTagMutabilityCmd)
}
