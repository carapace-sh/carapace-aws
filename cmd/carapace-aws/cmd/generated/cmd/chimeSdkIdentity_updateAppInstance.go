package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_updateAppInstanceCmd = &cobra.Command{
	Use:   "update-app-instance",
	Short: "Updates `AppInstance` metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_updateAppInstanceCmd).Standalone()

	chimeSdkIdentity_updateAppInstanceCmd.Flags().String("app-instance-arn", "", "The ARN of the `AppInstance`.")
	chimeSdkIdentity_updateAppInstanceCmd.Flags().String("metadata", "", "The metadata that you want to change.")
	chimeSdkIdentity_updateAppInstanceCmd.Flags().String("name", "", "The name that you want to change.")
	chimeSdkIdentity_updateAppInstanceCmd.MarkFlagRequired("app-instance-arn")
	chimeSdkIdentity_updateAppInstanceCmd.MarkFlagRequired("metadata")
	chimeSdkIdentity_updateAppInstanceCmd.MarkFlagRequired("name")
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_updateAppInstanceCmd)
}
