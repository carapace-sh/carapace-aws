package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_describeImageScanFindingsCmd = &cobra.Command{
	Use:   "describe-image-scan-findings",
	Short: "Returns the scan findings for the specified image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_describeImageScanFindingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_describeImageScanFindingsCmd).Standalone()

		ecr_describeImageScanFindingsCmd.Flags().String("image-id", "", "")
		ecr_describeImageScanFindingsCmd.Flags().String("max-results", "", "The maximum number of image scan results returned by `DescribeImageScanFindings` in paginated output.")
		ecr_describeImageScanFindingsCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `DescribeImageScanFindings` request where `maxResults` was used and the results exceeded the value of that parameter.")
		ecr_describeImageScanFindingsCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry that contains the repository in which to describe the image scan findings for.")
		ecr_describeImageScanFindingsCmd.Flags().String("repository-name", "", "The repository for the image for which to describe the scan findings.")
		ecr_describeImageScanFindingsCmd.MarkFlagRequired("image-id")
		ecr_describeImageScanFindingsCmd.MarkFlagRequired("repository-name")
	})
	ecrCmd.AddCommand(ecr_describeImageScanFindingsCmd)
}
