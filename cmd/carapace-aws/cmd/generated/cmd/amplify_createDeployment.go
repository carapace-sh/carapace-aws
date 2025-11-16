package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_createDeploymentCmd = &cobra.Command{
	Use:   "create-deployment",
	Short: "Creates a deployment for a manually deployed Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_createDeploymentCmd).Standalone()

	amplify_createDeploymentCmd.Flags().String("app-id", "", "The unique ID for an Amplify app.")
	amplify_createDeploymentCmd.Flags().String("branch-name", "", "The name of the branch to use for the job.")
	amplify_createDeploymentCmd.Flags().String("file-map", "", "An optional file map that contains the file name as the key and the file content md5 hash as the value.")
	amplify_createDeploymentCmd.MarkFlagRequired("app-id")
	amplify_createDeploymentCmd.MarkFlagRequired("branch-name")
	amplifyCmd.AddCommand(amplify_createDeploymentCmd)
}
