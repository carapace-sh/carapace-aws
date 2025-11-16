package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_getPersistentAppUipresignedUrlCmd = &cobra.Command{
	Use:   "get-persistent-app-uipresigned-url",
	Short: "The presigned URL properties for the cluster's application user interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_getPersistentAppUipresignedUrlCmd).Standalone()

	emr_getPersistentAppUipresignedUrlCmd.Flags().String("application-id", "", "The application ID associated with the presigned URL.")
	emr_getPersistentAppUipresignedUrlCmd.Flags().String("auth-proxy-call", "", "A boolean that represents if the caller is an authentication proxy call.")
	emr_getPersistentAppUipresignedUrlCmd.Flags().String("execution-role-arn", "", "The execution role ARN associated with the presigned URL.")
	emr_getPersistentAppUipresignedUrlCmd.Flags().String("persistent-app-uiid", "", "The persistent application user interface ID associated with the presigned URL.")
	emr_getPersistentAppUipresignedUrlCmd.Flags().String("persistent-app-uitype", "", "The persistent application user interface type associated with the presigned URL.")
	emr_getPersistentAppUipresignedUrlCmd.MarkFlagRequired("persistent-app-uiid")
	emrCmd.AddCommand(emr_getPersistentAppUipresignedUrlCmd)
}
