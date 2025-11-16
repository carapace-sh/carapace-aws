package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_listServiceVersionsCmd = &cobra.Command{
	Use:   "list-service-versions",
	Short: "Lists all supported versions for Snow on-device services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_listServiceVersionsCmd).Standalone()

	snowball_listServiceVersionsCmd.Flags().String("dependent-services", "", "A list of names and versions of dependant services of the requested service.")
	snowball_listServiceVersionsCmd.Flags().String("max-results", "", "The maximum number of `ListServiceVersions` objects to return.")
	snowball_listServiceVersionsCmd.Flags().String("next-token", "", "Because HTTP requests are stateless, this is the starting point for the next list of returned `ListServiceVersionsRequest` versions.")
	snowball_listServiceVersionsCmd.Flags().String("service-name", "", "The name of the service for which you're requesting supported versions.")
	snowball_listServiceVersionsCmd.MarkFlagRequired("service-name")
	snowballCmd.AddCommand(snowball_listServiceVersionsCmd)
}
