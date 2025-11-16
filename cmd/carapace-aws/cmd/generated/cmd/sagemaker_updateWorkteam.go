package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateWorkteamCmd = &cobra.Command{
	Use:   "update-workteam",
	Short: "Updates an existing work team with new member definitions or description.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateWorkteamCmd).Standalone()

	sagemaker_updateWorkteamCmd.Flags().String("description", "", "An updated description for the work team.")
	sagemaker_updateWorkteamCmd.Flags().String("member-definitions", "", "A list of `MemberDefinition` objects that contains objects that identify the workers that make up the work team.")
	sagemaker_updateWorkteamCmd.Flags().String("notification-configuration", "", "Configures SNS topic notifications for available or expiring work items")
	sagemaker_updateWorkteamCmd.Flags().String("worker-access-configuration", "", "Use this optional parameter to constrain access to an Amazon S3 resource based on the IP address using supported IAM global condition keys.")
	sagemaker_updateWorkteamCmd.Flags().String("workteam-name", "", "The name of the work team to update.")
	sagemaker_updateWorkteamCmd.MarkFlagRequired("workteam-name")
	sagemakerCmd.AddCommand(sagemaker_updateWorkteamCmd)
}
