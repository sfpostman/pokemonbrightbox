package root

import (
	"errors"
	e2elibrarysdk "example.com/e2e-library/sdk"
	"example.com/e2e-library/sdk/e2elibrarysdkconfig"
	"fmt"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "e2e-library",
	Short:        "e2e-library CLI",
	SilenceUsage: true,
}

var authConfigurators []func(client any)

func Execute() {
	cc.Init(&cc.Config{
		RootCmd:  rootCmd,
		Headings: cc.HiYellow + cc.Bold + cc.Underline,
		Commands: cc.HiGreen + cc.Bold,
		Example:  cc.Italic,
		ExecName: cc.Bold,
		Flags:    cc.Bold,
	})
	if err := rootCmd.Execute(); err != nil {
		printResponseBody(err)
		os.Exit(1)
	}
}

func printResponseBody(err error) {
	type bodyGetter interface{ GetBody() []byte }
	var e bodyGetter
	if errors.As(err, &e) {
		if body := e.GetBody(); len(body) > 0 {
			fmt.Fprintln(os.Stderr, string(body))
		}
	}
}

func AddCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

func RegisterAuthConfigurator(fn func(client any)) {
	authConfigurators = append(authConfigurators, fn)
}

func CreateSdkClient() *e2elibrarysdk.E2eLibrarySDK {
	sdkConfig := e2elibrarysdkconfig.NewConfig()

	if baseUrl := viper.GetString("base_url"); baseUrl != "" {
		sdkConfig.SetBaseURL(baseUrl)
	}

	client := e2elibrarysdk.NewE2eLibrarySDK(sdkConfig)

	// Apply credentials file values (lower priority than env vars / config file).
	for _, configure := range authConfigurators {
		configure(client)
	}

	// Re-apply viper values last so env vars and config file always win over
	// any values loaded from the credentials file above.
	if apiKey := viper.GetString("api_key"); apiKey != "" {
		client.SetAPIKey(apiKey)
	}
	return client
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.SetEnvPrefix("E2E_LIBRARY")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Fprintf(os.Stderr, "Warning: error reading config file: %v\n", err)
		}
	}
}
