package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listPartnerAppsCmd = &cobra.Command{
	Use:   "list-partner-apps",
	Short: "Lists all of the SageMaker Partner AI Apps in an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listPartnerAppsCmd).Standalone()

	sagemaker_listPartnerAppsCmd.Flags().String("max-results", "", "This parameter defines the maximum number of results that can be returned in a single response.")
	sagemaker_listPartnerAppsCmd.Flags().String("next-token", "", "If the previous response was truncated, you will receive this token.")
	sagemakerCmd.AddCommand(sagemaker_listPartnerAppsCmd)
}
