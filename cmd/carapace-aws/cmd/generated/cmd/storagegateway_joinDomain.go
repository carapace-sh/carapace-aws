package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_joinDomainCmd = &cobra.Command{
	Use:   "join-domain",
	Short: "Adds a file gateway to an Active Directory domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_joinDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_joinDomainCmd).Standalone()

		storagegateway_joinDomainCmd.Flags().String("domain-controllers", "", "List of IP addresses, NetBIOS names, or host names of your domain server.")
		storagegateway_joinDomainCmd.Flags().String("domain-name", "", "The name of the domain that you want the gateway to join.")
		storagegateway_joinDomainCmd.Flags().String("gateway-arn", "", "The Amazon Resource Name (ARN) of the gateway.")
		storagegateway_joinDomainCmd.Flags().String("organizational-unit", "", "The organizational unit (OU) is a container in an Active Directory that can hold users, groups, computers, and other OUs and this parameter specifies the OU that the gateway will join within the AD domain.")
		storagegateway_joinDomainCmd.Flags().String("password", "", "Sets the password of the user who has permission to add the gateway to the Active Directory domain.")
		storagegateway_joinDomainCmd.Flags().String("timeout-in-seconds", "", "Specifies the time in seconds, in which the `JoinDomain` operation must complete.")
		storagegateway_joinDomainCmd.Flags().String("user-name", "", "Sets the user name of user who has permission to add the gateway to the Active Directory domain.")
		storagegateway_joinDomainCmd.MarkFlagRequired("domain-name")
		storagegateway_joinDomainCmd.MarkFlagRequired("gateway-arn")
		storagegateway_joinDomainCmd.MarkFlagRequired("password")
		storagegateway_joinDomainCmd.MarkFlagRequired("user-name")
	})
	storagegatewayCmd.AddCommand(storagegateway_joinDomainCmd)
}
