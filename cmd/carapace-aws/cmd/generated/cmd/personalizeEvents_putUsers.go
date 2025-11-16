package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalizeEvents_putUsersCmd = &cobra.Command{
	Use:   "put-users",
	Short: "Adds one or more users to a Users dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalizeEvents_putUsersCmd).Standalone()

	personalizeEvents_putUsersCmd.Flags().String("dataset-arn", "", "The Amazon Resource Name (ARN) of the Users dataset you are adding the user or users to.")
	personalizeEvents_putUsersCmd.Flags().String("users", "", "A list of user data.")
	personalizeEvents_putUsersCmd.MarkFlagRequired("dataset-arn")
	personalizeEvents_putUsersCmd.MarkFlagRequired("users")
	personalizeEventsCmd.AddCommand(personalizeEvents_putUsersCmd)
}
