package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_listResourceProfileDetectionsCmd = &cobra.Command{
	Use:   "list-resource-profile-detections",
	Short: "Retrieves information about the types and amount of sensitive data that Amazon Macie found in an S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_listResourceProfileDetectionsCmd).Standalone()

	macie2_listResourceProfileDetectionsCmd.Flags().String("max-results", "", "The maximum number of items to include in each page of a paginated response.")
	macie2_listResourceProfileDetectionsCmd.Flags().String("next-token", "", "The nextToken string that specifies which page of results to return in a paginated response.")
	macie2_listResourceProfileDetectionsCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the S3 bucket that the request applies to.")
	macie2_listResourceProfileDetectionsCmd.MarkFlagRequired("resource-arn")
	macie2Cmd.AddCommand(macie2_listResourceProfileDetectionsCmd)
}
