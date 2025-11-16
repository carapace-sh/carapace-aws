package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_updateHealthCheckCmd = &cobra.Command{
	Use:   "update-health-check",
	Short: "Updates an existing health check.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_updateHealthCheckCmd).Standalone()

	route53_updateHealthCheckCmd.Flags().String("alarm-identifier", "", "A complex type that identifies the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether the specified health check is healthy.")
	route53_updateHealthCheckCmd.Flags().String("child-health-checks", "", "A complex type that contains one `ChildHealthCheck` element for each health check that you want to associate with a `CALCULATED` health check.")
	route53_updateHealthCheckCmd.Flags().String("disabled", "", "Stops Route 53 from performing health checks.")
	route53_updateHealthCheckCmd.Flags().String("enable-sni", "", "Specify whether you want Amazon Route 53 to send the value of `FullyQualifiedDomainName` to the endpoint in the `client_hello` message during `TLS` negotiation.")
	route53_updateHealthCheckCmd.Flags().String("failure-threshold", "", "The number of consecutive health checks that an endpoint must pass or fail for Amazon Route 53 to change the current status of the endpoint from unhealthy to healthy or vice versa.")
	route53_updateHealthCheckCmd.Flags().String("fully-qualified-domain-name", "", "Amazon Route 53 behavior depends on whether you specify a value for `IPAddress`.")
	route53_updateHealthCheckCmd.Flags().String("health-check-id", "", "The ID for the health check for which you want detailed information.")
	route53_updateHealthCheckCmd.Flags().String("health-check-version", "", "A sequential counter that Amazon Route 53 sets to `1` when you create a health check and increments by 1 each time you update settings for the health check.")
	route53_updateHealthCheckCmd.Flags().String("health-threshold", "", "The number of child health checks that are associated with a `CALCULATED` health that Amazon Route 53 must consider healthy for the `CALCULATED` health check to be considered healthy.")
	route53_updateHealthCheckCmd.Flags().String("insufficient-data-health-status", "", "When CloudWatch has insufficient data about the metric to determine the alarm state, the status that you want Amazon Route 53 to assign to the health check:")
	route53_updateHealthCheckCmd.Flags().String("inverted", "", "Specify whether you want Amazon Route 53 to invert the status of a health check, for example, to consider a health check unhealthy when it otherwise would be considered healthy.")
	route53_updateHealthCheckCmd.Flags().String("ipaddress", "", "The IPv4 or IPv6 IP address for the endpoint that you want Amazon Route 53 to perform health checks on.")
	route53_updateHealthCheckCmd.Flags().String("port", "", "The port on the endpoint that you want Amazon Route 53 to perform health checks on.")
	route53_updateHealthCheckCmd.Flags().String("regions", "", "A complex type that contains one `Region` element for each region that you want Amazon Route 53 health checkers to check the specified endpoint from.")
	route53_updateHealthCheckCmd.Flags().String("reset-elements", "", "A complex type that contains one `ResettableElementName` element for each element that you want to reset to the default value.")
	route53_updateHealthCheckCmd.Flags().String("resource-path", "", "The path that you want Amazon Route 53 to request when performing health checks.")
	route53_updateHealthCheckCmd.Flags().String("search-string", "", "If the value of `Type` is `HTTP_STR_MATCH` or `HTTPS_STR_MATCH`, the string that you want Amazon Route 53 to search for in the response body from the specified resource.")
	route53_updateHealthCheckCmd.MarkFlagRequired("health-check-id")
	route53Cmd.AddCommand(route53_updateHealthCheckCmd)
}
