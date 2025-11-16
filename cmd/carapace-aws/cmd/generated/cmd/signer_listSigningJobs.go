package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signer_listSigningJobsCmd = &cobra.Command{
	Use:   "list-signing-jobs",
	Short: "Lists all your signing jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signer_listSigningJobsCmd).Standalone()

	signer_listSigningJobsCmd.Flags().String("is-revoked", "", "Filters results to return only signing jobs with revoked signatures.")
	signer_listSigningJobsCmd.Flags().String("job-invoker", "", "Filters results to return only signing jobs initiated by a specified IAM entity.")
	signer_listSigningJobsCmd.Flags().String("max-results", "", "Specifies the maximum number of items to return in the response.")
	signer_listSigningJobsCmd.Flags().String("next-token", "", "String for specifying the next set of paginated results to return.")
	signer_listSigningJobsCmd.Flags().String("platform-id", "", "The ID of microcontroller platform that you specified for the distribution of your code image.")
	signer_listSigningJobsCmd.Flags().String("requested-by", "", "The IAM principal that requested the signing job.")
	signer_listSigningJobsCmd.Flags().String("signature-expires-after", "", "Filters results to return only signing jobs with signatures expiring after a specified timestamp.")
	signer_listSigningJobsCmd.Flags().String("signature-expires-before", "", "Filters results to return only signing jobs with signatures expiring before a specified timestamp.")
	signer_listSigningJobsCmd.Flags().String("status", "", "A status value with which to filter your results.")
	signerCmd.AddCommand(signer_listSigningJobsCmd)
}
