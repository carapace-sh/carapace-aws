package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_listResourceProfileArtifactsCmd = &cobra.Command{
	Use:   "list-resource-profile-artifacts",
	Short: "Retrieves information about objects that Amazon Macie selected from an S3 bucket for automated sensitive data discovery.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_listResourceProfileArtifactsCmd).Standalone()

	macie2_listResourceProfileArtifactsCmd.Flags().String("next-token", "", "The nextToken string that specifies which page of results to return in a paginated response.")
	macie2_listResourceProfileArtifactsCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the S3 bucket that the request applies to.")
	macie2_listResourceProfileArtifactsCmd.MarkFlagRequired("resource-arn")
	macie2Cmd.AddCommand(macie2_listResourceProfileArtifactsCmd)
}
