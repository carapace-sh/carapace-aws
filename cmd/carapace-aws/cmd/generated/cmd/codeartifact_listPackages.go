package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_listPackagesCmd = &cobra.Command{
	Use:   "list-packages",
	Short: "Returns a list of [PackageSummary](https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageSummary.html) objects for packages in a repository that match the request parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_listPackagesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeartifact_listPackagesCmd).Standalone()

		codeartifact_listPackagesCmd.Flags().String("domain", "", "The name of the domain that contains the repository that contains the requested packages.")
		codeartifact_listPackagesCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
		codeartifact_listPackagesCmd.Flags().String("format", "", "The format used to filter requested packages.")
		codeartifact_listPackagesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		codeartifact_listPackagesCmd.Flags().String("namespace", "", "The namespace prefix used to filter requested packages.")
		codeartifact_listPackagesCmd.Flags().String("next-token", "", "The token for the next set of results.")
		codeartifact_listPackagesCmd.Flags().String("package-prefix", "", "A prefix used to filter requested packages.")
		codeartifact_listPackagesCmd.Flags().String("publish", "", "The value of the `Publish` package origin control restriction used to filter requested packages.")
		codeartifact_listPackagesCmd.Flags().String("repository", "", "The name of the repository that contains the requested packages.")
		codeartifact_listPackagesCmd.Flags().String("upstream", "", "The value of the `Upstream` package origin control restriction used to filter requested packages.")
		codeartifact_listPackagesCmd.MarkFlagRequired("domain")
		codeartifact_listPackagesCmd.MarkFlagRequired("repository")
	})
	codeartifactCmd.AddCommand(codeartifact_listPackagesCmd)
}
