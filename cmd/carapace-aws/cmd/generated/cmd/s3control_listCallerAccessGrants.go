package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_listCallerAccessGrantsCmd = &cobra.Command{
	Use:   "list-caller-access-grants",
	Short: "Use this API to list the access grants that grant the caller access to Amazon S3 data through S3 Access Grants.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_listCallerAccessGrantsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_listCallerAccessGrantsCmd).Standalone()

		s3control_listCallerAccessGrantsCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the S3 Access Grants instance.")
		s3control_listCallerAccessGrantsCmd.Flags().Bool("allowed-by-application", false, "If this optional parameter is passed in the request, a filter is applied to the results.")
		s3control_listCallerAccessGrantsCmd.Flags().String("grant-scope", "", "The S3 path of the data that you would like to access.")
		s3control_listCallerAccessGrantsCmd.Flags().String("max-results", "", "The maximum number of access grants that you would like returned in the `List Caller Access Grants` response.")
		s3control_listCallerAccessGrantsCmd.Flags().String("next-token", "", "A pagination token to request the next page of results.")
		s3control_listCallerAccessGrantsCmd.Flags().Bool("no-allowed-by-application", false, "If this optional parameter is passed in the request, a filter is applied to the results.")
		s3control_listCallerAccessGrantsCmd.MarkFlagRequired("account-id")
		s3control_listCallerAccessGrantsCmd.Flag("no-allowed-by-application").Hidden = true
	})
	s3controlCmd.AddCommand(s3control_listCallerAccessGrantsCmd)
}
