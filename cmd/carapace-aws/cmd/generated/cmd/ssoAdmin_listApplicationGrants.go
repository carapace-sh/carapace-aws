package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_listApplicationGrantsCmd = &cobra.Command{
	Use:   "list-application-grants",
	Short: "List the grants associated with an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_listApplicationGrantsCmd).Standalone()

	ssoAdmin_listApplicationGrantsCmd.Flags().String("application-arn", "", "Specifies the ARN of the application whose grants you want to list.")
	ssoAdmin_listApplicationGrantsCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
	ssoAdmin_listApplicationGrantsCmd.MarkFlagRequired("application-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_listApplicationGrantsCmd)
}
