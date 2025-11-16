package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_listAccessGrantsInstancesCmd = &cobra.Command{
	Use:   "list-access-grants-instances",
	Short: "Returns a list of S3 Access Grants instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_listAccessGrantsInstancesCmd).Standalone()

	s3control_listAccessGrantsInstancesCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the S3 Access Grants instance.")
	s3control_listAccessGrantsInstancesCmd.Flags().String("max-results", "", "The maximum number of access grants that you would like returned in the `List Access Grants` response.")
	s3control_listAccessGrantsInstancesCmd.Flags().String("next-token", "", "A pagination token to request the next page of results.")
	s3control_listAccessGrantsInstancesCmd.MarkFlagRequired("account-id")
	s3controlCmd.AddCommand(s3control_listAccessGrantsInstancesCmd)
}
