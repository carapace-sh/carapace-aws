package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_listPackageVersionsCmd = &cobra.Command{
	Use:   "list-package-versions",
	Short: "Returns a list of [PackageVersionSummary](https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageVersionSummary.html) objects for package versions in a repository that match the request parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_listPackageVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeartifact_listPackageVersionsCmd).Standalone()

		codeartifact_listPackageVersionsCmd.Flags().String("domain", "", "The name of the domain that contains the repository that contains the requested package versions.")
		codeartifact_listPackageVersionsCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
		codeartifact_listPackageVersionsCmd.Flags().String("format", "", "The format of the package versions you want to list.")
		codeartifact_listPackageVersionsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		codeartifact_listPackageVersionsCmd.Flags().String("namespace", "", "The namespace of the package that contains the requested package versions.")
		codeartifact_listPackageVersionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		codeartifact_listPackageVersionsCmd.Flags().String("origin-type", "", "The `originType` used to filter package versions.")
		codeartifact_listPackageVersionsCmd.Flags().String("package", "", "The name of the package for which you want to request package versions.")
		codeartifact_listPackageVersionsCmd.Flags().String("repository", "", "The name of the repository that contains the requested package versions.")
		codeartifact_listPackageVersionsCmd.Flags().String("sort-by", "", "How to sort the requested list of package versions.")
		codeartifact_listPackageVersionsCmd.Flags().String("status", "", "A string that filters the requested package versions by status.")
		codeartifact_listPackageVersionsCmd.MarkFlagRequired("domain")
		codeartifact_listPackageVersionsCmd.MarkFlagRequired("format")
		codeartifact_listPackageVersionsCmd.MarkFlagRequired("package")
		codeartifact_listPackageVersionsCmd.MarkFlagRequired("repository")
	})
	codeartifactCmd.AddCommand(codeartifact_listPackageVersionsCmd)
}
