package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listServiceSpecificCredentialsCmd = &cobra.Command{
	Use:   "list-service-specific-credentials",
	Short: "Returns information about the service-specific credentials associated with the specified IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listServiceSpecificCredentialsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_listServiceSpecificCredentialsCmd).Standalone()

		iam_listServiceSpecificCredentialsCmd.Flags().String("all-users", "", "A flag indicating whether to list service specific credentials for all users.")
		iam_listServiceSpecificCredentialsCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
		iam_listServiceSpecificCredentialsCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
		iam_listServiceSpecificCredentialsCmd.Flags().String("service-name", "", "Filters the returned results to only those for the specified Amazon Web Services service.")
		iam_listServiceSpecificCredentialsCmd.Flags().String("user-name", "", "The name of the user whose service-specific credentials you want information about.")
	})
	iamCmd.AddCommand(iam_listServiceSpecificCredentialsCmd)
}
