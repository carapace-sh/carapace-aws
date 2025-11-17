package cmd

import (
	"fmt"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-aws/cmd/carapace-aws/cmd/botocore"
	"github.com/carapace-sh/carapace-aws/cmd/carapace-aws/cmd/common"
	spec "github.com/carapace-sh/carapace-spec"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "aws",
	Short: "An enriched aws completer",
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	rootCmd.SetUsageFunc(func(c *cobra.Command) error { return nil })

	rootCmd.PersistentFlags().Bool("ca-bundle", false, "The CA certificate bundle to use when verifying SSL certificates.")
	rootCmd.PersistentFlags().Bool("cli-auto-prompt", false, "Automatically prompt for CLI input parameters.")
	rootCmd.PersistentFlags().String("cli-binary-format", "", "The formatting style to be used for binary blobs.")
	rootCmd.PersistentFlags().String("cli-connect-timeout", "", "The maximum socket connect time in seconds.")
	rootCmd.PersistentFlags().String("cli-read-timeout", "", "The maximum socket read time in seconds.")
	rootCmd.PersistentFlags().String("color", "", "Turn on/off color output.")
	rootCmd.PersistentFlags().Bool("debug", false, "Turn on debug logging.")
	rootCmd.PersistentFlags().String("endpoint-url", "", "Override command's default URL with the given URL.")
	rootCmd.PersistentFlags().Bool("no-cli-auto-prompt", false, "Disable automatically prompt for CLI input parameters.")
	rootCmd.PersistentFlags().Bool("no-cli-pager", false, "Disable cli pager for output.")
	rootCmd.PersistentFlags().Bool("no-paginate", false, "Disable automatic pagination.")
	rootCmd.PersistentFlags().Bool("no-sign-request", false, "Do not sign requests.")
	rootCmd.PersistentFlags().Bool("no-verify-ssl", false, "Override the default behavior of verifying SSL certificates.")
	rootCmd.PersistentFlags().String("output", "", "The formatting style for command output.")
	rootCmd.PersistentFlags().String("profile", "", "Use a specific profile from your credential file.")
	rootCmd.PersistentFlags().String("query", "", "A JMESPath query to use in filtering the response data.")
	rootCmd.PersistentFlags().String("region", "", "The region to use. Overrides config/env settings.")
	rootCmd.PersistentFlags().Bool("version", false, "Display the version of this tool.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cli-binary-format": carapace.ActionValues(
			"base64",
			"raw-in-base64-out",
		),
		"color":   carapace.ActionValues("on", "off", "auto").StyleF(style.ForKeyword),
		"output":  carapace.ActionValues("json", "text", "table", "yaml", "yaml-stream"),
		"profile": spec.ActionMacro("$carapace.tools.aws.Profiles"),
		"region":  spec.ActionMacro("$carapace.tools.aws.Regions"),
	})

	for name, description := range botocore.Services() {
		serviceCmd := &cobra.Command{
			Use:   name,
			Short: description,
			Run:   func(cmd *cobra.Command, args []string) {},
		}
		rootCmd.AddCommand(serviceCmd)
		carapace.Gen(serviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
			for name := range botocore.Operations(cmd.Use) {
				botoCommand, err := botocore.Get(fmt.Sprintf("aws.%s.%s.yaml", serviceCmd.Use, name))
				if err != nil {
					carapace.LOG.Println(err.Error()) // TODO handle error
					return
				}
				operationCmd := botoCommand.ToCobra()
				serviceCmd.AddCommand(operationCmd)

				carapace.Gen(operationCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
					if flag != nil && flag.Value.Type() == "string" {
						if _, ok := botoCommand.Completion.Flag[flag.Name]; !ok {
							return common.ActionBridgeAwsCompleter()
						}
					}
					return action
				})
			}
		})
	}

	for _, extension := range []string{
		"apptest",
		"cli-dev",
		"configservice",
		"configure",
		"ddb",
		"deploy",
		"history",
		"iotfleethub",
		"lookoutmetrics",
		"lookoutvision",
		"qldb",
		"qldb-session",
		"robomaker",
		"s3api",
		"sms",
	} {
		subCmd := &cobra.Command{
			Use:                extension,
			Run:                func(cmd *cobra.Command, args []string) {},
			DisableFlagParsing: true,
		}
		rootCmd.AddCommand(subCmd)

		carapace.Gen(subCmd).PositionalAnyCompletion(
			common.ActionBridgeAwsCompleter(),
		)
	}
}
