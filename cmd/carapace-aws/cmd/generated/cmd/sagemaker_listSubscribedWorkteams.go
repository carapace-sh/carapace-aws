package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listSubscribedWorkteamsCmd = &cobra.Command{
	Use:   "list-subscribed-workteams",
	Short: "Gets a list of the work teams that you are subscribed to in the Amazon Web Services Marketplace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listSubscribedWorkteamsCmd).Standalone()

	sagemaker_listSubscribedWorkteamsCmd.Flags().String("max-results", "", "The maximum number of work teams to return in each page of the response.")
	sagemaker_listSubscribedWorkteamsCmd.Flags().String("name-contains", "", "A string in the work team name.")
	sagemaker_listSubscribedWorkteamsCmd.Flags().String("next-token", "", "If the result of the previous `ListSubscribedWorkteams` request was truncated, the response includes a `NextToken`.")
	sagemakerCmd.AddCommand(sagemaker_listSubscribedWorkteamsCmd)
}
