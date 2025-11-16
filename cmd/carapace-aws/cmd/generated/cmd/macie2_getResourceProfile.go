package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_getResourceProfileCmd = &cobra.Command{
	Use:   "get-resource-profile",
	Short: "Retrieves (queries) sensitive data discovery statistics and the sensitivity score for an S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_getResourceProfileCmd).Standalone()

	macie2_getResourceProfileCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the S3 bucket that the request applies to.")
	macie2_getResourceProfileCmd.MarkFlagRequired("resource-arn")
	macie2Cmd.AddCommand(macie2_getResourceProfileCmd)
}
