package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrPublic_describeRegistriesCmd = &cobra.Command{
	Use:   "describe-registries",
	Short: "Returns details for a public registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrPublic_describeRegistriesCmd).Standalone()

	ecrPublic_describeRegistriesCmd.Flags().String("max-results", "", "The maximum number of repository results that's returned by `DescribeRegistries` in paginated output.")
	ecrPublic_describeRegistriesCmd.Flags().String("next-token", "", "The `nextToken` value that's returned from a previous paginated `DescribeRegistries` request where `maxResults` was used and the results exceeded the value of that parameter.")
	ecrPublicCmd.AddCommand(ecrPublic_describeRegistriesCmd)
}
