/*
 * lutymaps - a mapping engine for luty
 *
 * Copyright (c) 2023 2023 Michael D Henderson
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package cli

import (
	"errors"
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	ENV_PREFIX = "LUTYMAPS"
)

var (
	cliConfig Config
)

// cmdMain represents the base command when called without any subcommands
var cmdMain = &cobra.Command{
	Use:   "lutymaps",
	Short: "Luty's Personal Map engine",
	Long:  `Lutymaps is a map engine for Luty.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// bind viper and cobra here since this hook runs early and always
		return bindConfig(cmd)
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("env: %-30s == %q\n", "HOME", cliConfig.HomeFolder)
		log.Printf("env: %-30s == %q\n", ENV_PREFIX+"_CONFIG", viper.ConfigFileUsed())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(cmdMain.Execute())
}

func init() {
	cmdMain.PersistentFlags().StringVar(&cliConfig.ConfigFile, "config", "", "config file (default is ~/."+strings.ToLower(ENV_PREFIX)+".json)")
	cmdMain.PersistentFlags().BoolVar(&cliConfig.Flags.Test, "test", false, "test mode")
	cmdMain.PersistentFlags().BoolVar(&cliConfig.Flags.Verbose, "verbose", false, "verbose mode")
}

// bindConfig reads in config file and ENV variables if set.
// logic for binding viper and cobra taken from
// https://carolynvanslyck.com/blog/2020/08/sting-of-the-viper/
func bindConfig(cmd *cobra.Command) error {
	var err error

	// Find home directory.
	cliConfig.HomeFolder, err = homedir.Dir()
	if err != nil {
		return err
	}

	if cliConfig.ConfigFile == "" { // use default location
		cliConfig.ConfigFile = filepath.Join(cliConfig.HomeFolder, "."+strings.ToLower(ENV_PREFIX)+".json")
	}
	viper.SetConfigFile(cliConfig.ConfigFile)

	// Try to read the config file. Ignore file-not-found errors.
	if err = viper.ReadInConfig(); err != nil {
		var errConfigFileNotFound viper.ConfigFileNotFoundError
		if errors.As(err, &errConfigFileNotFound) || os.IsNotExist(err) {
			// ignore file not found errors
		} else {
			return err
		}
	} else {
		log.Printf("viper: using config file: %q\n", viper.ConfigFileUsed())
	}
	// write out our configuration for debugging
	viperFilesPath, ok := viper.Get("files.path").(string)
	if !ok {
		log.Printf("viper: files.path is missing\n")
		viperFilesPath = "."
	}
	debugConfigFile := filepath.Join(viperFilesPath, "viper."+strings.ToLower(ENV_PREFIX)+".json")
	log.Printf("viper: writing debug config file: %q\n", debugConfigFile)
	if err = viper.WriteConfigAs(debugConfigFile); err != nil {
		return err
	}

	// read in environment variables that match
	viper.SetEnvPrefix(ENV_PREFIX)
	viper.AutomaticEnv()

	// bind the current command's flags to viper
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Environment variables can't have dashes in them, so bind them to their equivalent
		// keys with underscores, e.g. --favorite-color to STING_FAVORITE_COLOR
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			_ = viper.BindEnv(f.Name, fmt.Sprintf("%s_%s", ENV_PREFIX, envVarSuffix))
		}

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && viper.IsSet(f.Name) {
			val := viper.Get(f.Name)
			_ = cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})

	return nil
}
