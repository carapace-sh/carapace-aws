package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_getRepositoryEndpointCmd = &cobra.Command{
	Use:   "get-repository-endpoint",
	Short: "Returns the endpoint of a repository for a specific package format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_getRepositoryEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeartifact_getRepositoryEndpointCmd).Standalone()

		codeartifact_getRepositoryEndpointCmd.Flags().String("domain", "", "The name of the domain that contains the repository.")
		codeartifact_getRepositoryEndpointCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain that contains the repository.")
		codeartifact_getRepositoryEndpointCmd.Flags().String("endpoint-type", "", "A string that specifies the type of endpoint.")
		codeartifact_getRepositoryEndpointCmd.Flags().String("format", "", "Returns which endpoint of a repository to return.")
		codeartifact_getRepositoryEndpointCmd.Flags().String("repository", "", "The name of the repository.")
		codeartifact_getRepositoryEndpointCmd.MarkFlagRequired("domain")
		codeartifact_getRepositoryEndpointCmd.MarkFlagRequired("format")
		codeartifact_getRepositoryEndpointCmd.MarkFlagRequired("repository")
	})
	codeartifactCmd.AddCommand(codeartifact_getRepositoryEndpointCmd)
}
