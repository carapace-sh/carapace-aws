package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_createDomainNameCmd = &cobra.Command{
	Use:   "create-domain-name",
	Short: "Creates a custom `DomainName` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_createDomainNameCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_createDomainNameCmd).Standalone()

		appsync_createDomainNameCmd.Flags().String("certificate-arn", "", "The Amazon Resource Name (ARN) of the certificate.")
		appsync_createDomainNameCmd.Flags().String("description", "", "A description of the `DomainName`.")
		appsync_createDomainNameCmd.Flags().String("domain-name", "", "The domain name.")
		appsync_createDomainNameCmd.Flags().String("tags", "", "")
		appsync_createDomainNameCmd.MarkFlagRequired("certificate-arn")
		appsync_createDomainNameCmd.MarkFlagRequired("domain-name")
	})
	appsyncCmd.AddCommand(appsync_createDomainNameCmd)
}
