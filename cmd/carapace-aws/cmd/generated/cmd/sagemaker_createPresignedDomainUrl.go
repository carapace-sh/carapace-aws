package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createPresignedDomainUrlCmd = &cobra.Command{
	Use:   "create-presigned-domain-url",
	Short: "Creates a URL for a specified UserProfile in a Domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createPresignedDomainUrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createPresignedDomainUrlCmd).Standalone()

		sagemaker_createPresignedDomainUrlCmd.Flags().String("domain-id", "", "The domain ID.")
		sagemaker_createPresignedDomainUrlCmd.Flags().String("expires-in-seconds", "", "The number of seconds until the pre-signed URL expires.")
		sagemaker_createPresignedDomainUrlCmd.Flags().String("landing-uri", "", "The landing page that the user is directed to when accessing the presigned URL.")
		sagemaker_createPresignedDomainUrlCmd.Flags().String("session-expiration-duration-in-seconds", "", "The session expiration duration in seconds.")
		sagemaker_createPresignedDomainUrlCmd.Flags().String("space-name", "", "The name of the space.")
		sagemaker_createPresignedDomainUrlCmd.Flags().String("user-profile-name", "", "The name of the UserProfile to sign-in as.")
		sagemaker_createPresignedDomainUrlCmd.MarkFlagRequired("domain-id")
		sagemaker_createPresignedDomainUrlCmd.MarkFlagRequired("user-profile-name")
	})
	sagemakerCmd.AddCommand(sagemaker_createPresignedDomainUrlCmd)
}
