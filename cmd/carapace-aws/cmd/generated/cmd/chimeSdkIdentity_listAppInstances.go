package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_listAppInstancesCmd = &cobra.Command{
	Use:   "list-app-instances",
	Short: "Lists all Amazon Chime `AppInstance`s created under a single AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_listAppInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkIdentity_listAppInstancesCmd).Standalone()

		chimeSdkIdentity_listAppInstancesCmd.Flags().String("max-results", "", "The maximum number of `AppInstance`s that you want to return.")
		chimeSdkIdentity_listAppInstancesCmd.Flags().String("next-token", "", "The token passed by previous API requests until you reach the maximum number of `AppInstances`.")
	})
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_listAppInstancesCmd)
}
