package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed dsql/aws.dsql.generate-db-connect-admin-auth-token.yaml
var dsql_generateDbConnectAdminAuthToken []byte

//go:embed dsql/aws.dsql.generate-db-connect-auth-token.yaml
var dsql_generateDbConnectAuthToken []byte

func init() {
	customizations["dsql"] = func(cmd *command.Command) error {
		for _, spec := range [][]byte{
			dsql_generateDbConnectAdminAuthToken,
			dsql_generateDbConnectAuthToken,
		} {
			var specCommand command.Command
			if err := yaml.Unmarshal(spec, &specCommand); err != nil {
				return err
			}
			cmd.Commands = append(cmd.Commands, specCommand)
		}
		return nil
	}
}
