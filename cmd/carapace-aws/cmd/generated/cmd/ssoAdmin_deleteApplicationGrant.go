package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_deleteApplicationGrantCmd = &cobra.Command{
	Use:   "delete-application-grant",
	Short: "Deletes a grant from an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_deleteApplicationGrantCmd).Standalone()

	ssoAdmin_deleteApplicationGrantCmd.Flags().String("application-arn", "", "Specifies the ARN of the application with the grant to delete.")
	ssoAdmin_deleteApplicationGrantCmd.Flags().String("grant-type", "", "Specifies the type of grant to delete from the application.")
	ssoAdmin_deleteApplicationGrantCmd.MarkFlagRequired("application-arn")
	ssoAdmin_deleteApplicationGrantCmd.MarkFlagRequired("grant-type")
	ssoAdminCmd.AddCommand(ssoAdmin_deleteApplicationGrantCmd)
}
