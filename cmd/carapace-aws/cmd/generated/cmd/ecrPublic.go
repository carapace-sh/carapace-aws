package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrPublicCmd = &cobra.Command{
	Use:   "ecr-public",
	Short: "Amazon Elastic Container Registry Public\n\nAmazon Elastic Container Registry Public (Amazon ECR Public) is a managed container image registry service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrPublicCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecrPublicCmd).Standalone()

	})
	rootCmd.AddCommand(ecrPublicCmd)
}
