package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_listFargateProfilesCmd = &cobra.Command{
	Use:   "list-fargate-profiles",
	Short: "Lists the Fargate profiles associated with the specified cluster in your Amazon Web Services account in the specified Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_listFargateProfilesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_listFargateProfilesCmd).Standalone()

		eks_listFargateProfilesCmd.Flags().String("cluster-name", "", "The name of your cluster.")
		eks_listFargateProfilesCmd.Flags().String("max-results", "", "The maximum number of results, returned in paginated output.")
		eks_listFargateProfilesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated request, where `maxResults` was used and the results exceeded the value of that parameter.")
		eks_listFargateProfilesCmd.MarkFlagRequired("cluster-name")
	})
	eksCmd.AddCommand(eks_listFargateProfilesCmd)
}
