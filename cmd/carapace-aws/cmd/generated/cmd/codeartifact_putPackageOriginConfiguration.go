package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_putPackageOriginConfigurationCmd = &cobra.Command{
	Use:   "put-package-origin-configuration",
	Short: "Sets the package origin configuration for a package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_putPackageOriginConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeartifact_putPackageOriginConfigurationCmd).Standalone()

		codeartifact_putPackageOriginConfigurationCmd.Flags().String("domain", "", "The name of the domain that contains the repository that contains the package.")
		codeartifact_putPackageOriginConfigurationCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
		codeartifact_putPackageOriginConfigurationCmd.Flags().String("format", "", "A format that specifies the type of the package to be updated.")
		codeartifact_putPackageOriginConfigurationCmd.Flags().String("namespace", "", "The namespace of the package to be updated.")
		codeartifact_putPackageOriginConfigurationCmd.Flags().String("package", "", "The name of the package to be updated.")
		codeartifact_putPackageOriginConfigurationCmd.Flags().String("repository", "", "The name of the repository that contains the package.")
		codeartifact_putPackageOriginConfigurationCmd.Flags().String("restrictions", "", "A [PackageOriginRestrictions](https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageOriginRestrictions.html) object that contains information about the `upstream` and `publish` package origin restrictions.")
		codeartifact_putPackageOriginConfigurationCmd.MarkFlagRequired("domain")
		codeartifact_putPackageOriginConfigurationCmd.MarkFlagRequired("format")
		codeartifact_putPackageOriginConfigurationCmd.MarkFlagRequired("package")
		codeartifact_putPackageOriginConfigurationCmd.MarkFlagRequired("repository")
		codeartifact_putPackageOriginConfigurationCmd.MarkFlagRequired("restrictions")
	})
	codeartifactCmd.AddCommand(codeartifact_putPackageOriginConfigurationCmd)
}
