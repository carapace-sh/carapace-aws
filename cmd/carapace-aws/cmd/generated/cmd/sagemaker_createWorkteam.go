package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createWorkteamCmd = &cobra.Command{
	Use:   "create-workteam",
	Short: "Creates a new work team for labeling your data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createWorkteamCmd).Standalone()

	sagemaker_createWorkteamCmd.Flags().String("description", "", "A description of the work team.")
	sagemaker_createWorkteamCmd.Flags().String("member-definitions", "", "A list of `MemberDefinition` objects that contains objects that identify the workers that make up the work team.")
	sagemaker_createWorkteamCmd.Flags().String("notification-configuration", "", "Configures notification of workers regarding available or expiring work items.")
	sagemaker_createWorkteamCmd.Flags().String("tags", "", "An array of key-value pairs.")
	sagemaker_createWorkteamCmd.Flags().String("worker-access-configuration", "", "Use this optional parameter to constrain access to an Amazon S3 resource based on the IP address using supported IAM global condition keys.")
	sagemaker_createWorkteamCmd.Flags().String("workforce-name", "", "The name of the workforce.")
	sagemaker_createWorkteamCmd.Flags().String("workteam-name", "", "The name of the work team.")
	sagemaker_createWorkteamCmd.MarkFlagRequired("description")
	sagemaker_createWorkteamCmd.MarkFlagRequired("member-definitions")
	sagemaker_createWorkteamCmd.MarkFlagRequired("workteam-name")
	sagemakerCmd.AddCommand(sagemaker_createWorkteamCmd)
}
