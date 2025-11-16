package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_listCustomDomainAssociationsCmd = &cobra.Command{
	Use:   "list-custom-domain-associations",
	Short: "Lists custom domain associations for Amazon Redshift Serverless.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_listCustomDomainAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_listCustomDomainAssociationsCmd).Standalone()

		redshiftServerless_listCustomDomainAssociationsCmd.Flags().String("custom-domain-certificate-arn", "", "The custom domain name’s certificate Amazon resource name (ARN).")
		redshiftServerless_listCustomDomainAssociationsCmd.Flags().String("custom-domain-name", "", "The custom domain name associated with the workgroup.")
		redshiftServerless_listCustomDomainAssociationsCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
		redshiftServerless_listCustomDomainAssociationsCmd.Flags().String("next-token", "", "When `nextToken` is returned, there are more results available.")
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_listCustomDomainAssociationsCmd)
}
