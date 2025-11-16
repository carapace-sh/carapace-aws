package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_describeAddonVersionsCmd = &cobra.Command{
	Use:   "describe-addon-versions",
	Short: "Describes the versions for an add-on.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_describeAddonVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_describeAddonVersionsCmd).Standalone()

		eks_describeAddonVersionsCmd.Flags().String("addon-name", "", "The name of the add-on.")
		eks_describeAddonVersionsCmd.Flags().String("kubernetes-version", "", "The Kubernetes versions that you can use the add-on with.")
		eks_describeAddonVersionsCmd.Flags().String("max-results", "", "The maximum number of results, returned in paginated output.")
		eks_describeAddonVersionsCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated request, where `maxResults` was used and the results exceeded the value of that parameter.")
		eks_describeAddonVersionsCmd.Flags().String("owners", "", "The owner of the add-on.")
		eks_describeAddonVersionsCmd.Flags().String("publishers", "", "The publisher of the add-on.")
		eks_describeAddonVersionsCmd.Flags().String("types", "", "The type of the add-on.")
	})
	eksCmd.AddCommand(eks_describeAddonVersionsCmd)
}
