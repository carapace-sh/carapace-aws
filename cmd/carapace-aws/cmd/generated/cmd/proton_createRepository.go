package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_createRepositoryCmd = &cobra.Command{
	Use:   "create-repository",
	Short: "Create and register a link to a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_createRepositoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_createRepositoryCmd).Standalone()

		proton_createRepositoryCmd.Flags().String("connection-arn", "", "The Amazon Resource Name (ARN) of your AWS CodeStar connection that connects Proton to your repository provider account.")
		proton_createRepositoryCmd.Flags().String("encryption-key", "", "The ARN of your customer Amazon Web Services Key Management Service (Amazon Web Services KMS) key.")
		proton_createRepositoryCmd.Flags().String("name", "", "The repository name (for example, `myrepos/myrepo`).")
		proton_createRepositoryCmd.Flags().String("provider", "", "The repository provider.")
		proton_createRepositoryCmd.Flags().String("tags", "", "An optional list of metadata items that you can associate with the Proton repository.")
		proton_createRepositoryCmd.MarkFlagRequired("connection-arn")
		proton_createRepositoryCmd.MarkFlagRequired("name")
		proton_createRepositoryCmd.MarkFlagRequired("provider")
	})
	protonCmd.AddCommand(proton_createRepositoryCmd)
}
