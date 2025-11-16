package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_updatePackageGroupOriginConfigurationCmd = &cobra.Command{
	Use:   "update-package-group-origin-configuration",
	Short: "Updates the package origin configuration for a package group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_updatePackageGroupOriginConfigurationCmd).Standalone()

	codeartifact_updatePackageGroupOriginConfigurationCmd.Flags().String("add-allowed-repositories", "", "The repository name and restrictions to add to the allowed repository list of the specified package group.")
	codeartifact_updatePackageGroupOriginConfigurationCmd.Flags().String("domain", "", "The name of the domain which contains the package group for which to update the origin configuration.")
	codeartifact_updatePackageGroupOriginConfigurationCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
	codeartifact_updatePackageGroupOriginConfigurationCmd.Flags().String("package-group", "", "The pattern of the package group for which to update the origin configuration.")
	codeartifact_updatePackageGroupOriginConfigurationCmd.Flags().String("remove-allowed-repositories", "", "The repository name and restrictions to remove from the allowed repository list of the specified package group.")
	codeartifact_updatePackageGroupOriginConfigurationCmd.Flags().String("restrictions", "", "The origin configuration settings that determine how package versions can enter repositories.")
	codeartifact_updatePackageGroupOriginConfigurationCmd.MarkFlagRequired("domain")
	codeartifact_updatePackageGroupOriginConfigurationCmd.MarkFlagRequired("package-group")
	codeartifactCmd.AddCommand(codeartifact_updatePackageGroupOriginConfigurationCmd)
}
