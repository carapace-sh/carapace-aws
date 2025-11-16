package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_testDnsanswerCmd = &cobra.Command{
	Use:   "test-dnsanswer",
	Short: "Gets the value that Amazon Route 53 returns in response to a DNS request for a specified record name and type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_testDnsanswerCmd).Standalone()

	route53_testDnsanswerCmd.Flags().String("edns0-client-subnet-ip", "", "If the resolver that you specified for resolverip supports EDNS0, specify the IPv4 or IPv6 address of a client in the applicable location, for example, `192.0.2.44` or `2001:db8:85a3::8a2e:370:7334`.")
	route53_testDnsanswerCmd.Flags().String("edns0-client-subnet-mask", "", "If you specify an IP address for `edns0clientsubnetip`, you can optionally specify the number of bits of the IP address that you want the checking tool to include in the DNS query.")
	route53_testDnsanswerCmd.Flags().String("hosted-zone-id", "", "The ID of the hosted zone that you want Amazon Route 53 to simulate a query for.")
	route53_testDnsanswerCmd.Flags().String("record-name", "", "The name of the resource record set that you want Amazon Route 53 to simulate a query for.")
	route53_testDnsanswerCmd.Flags().String("record-type", "", "The type of the resource record set.")
	route53_testDnsanswerCmd.Flags().String("resolver-ip", "", "If you want to simulate a request from a specific DNS resolver, specify the IP address for that resolver.")
	route53_testDnsanswerCmd.MarkFlagRequired("hosted-zone-id")
	route53_testDnsanswerCmd.MarkFlagRequired("record-name")
	route53_testDnsanswerCmd.MarkFlagRequired("record-type")
	route53Cmd.AddCommand(route53_testDnsanswerCmd)
}
