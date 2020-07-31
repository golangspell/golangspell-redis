package usecase

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golangspell/golangspell-redis/appcontext"
	"github.com/golangspell/golangspell-redis/domain"
	toolconfig "github.com/golangspell/golangspell/config"
	tooldomain "github.com/golangspell/golangspell/domain"
)

func addComponentConstantToContext(currentPath string, constantDefinition string) error {
	filePath := fmt.Sprintf("%s%sappcontext%scontext.go", currentPath, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	renderer := domain.GetRenderer()
	err := renderer.BackupExistingCode(filePath)
	if err != nil {
		return fmt.Errorf("An error occurred while trying to backup the context file. Error: %s", err.Error())
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("An error occurred while trying to read the context file. Error: %s", err.Error())
	}
	code := strings.ReplaceAll(
		string(content),
		"const (\n",
		fmt.Sprintf("const (\n%s\n", constantDefinition))
	err = ioutil.WriteFile(filePath, []byte(code), 0644)
	if err != nil {
		return fmt.Errorf("An error occurred while trying to update the context file. Error: %s", err.Error())
	}

	return nil
}

func addImportToMain(currentPath string, importPath string) error {
	filePath := fmt.Sprintf("%s%smain.go", currentPath, toolconfig.PlatformSeparator)
	renderer := domain.GetRenderer()
	err := renderer.BackupExistingCode(filePath)
	if err != nil {
		return fmt.Errorf("An error occurred while trying to backup the main file. Error: %s", err.Error())
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("An error occurred while trying to read the main file. Error: %s", err.Error())
	}

	if strings.Contains(string(content), importPath) {
		return nil
	}

	code := strings.ReplaceAll(
		string(content),
		"/config\"\n",
		fmt.Sprintf("/config\"\n_ \"%s\"\n", importPath))
	err = ioutil.WriteFile(filePath, []byte(code), 0644)
	if err != nil {
		return fmt.Errorf("An error occurred while trying to update the main file. Error: %s", err.Error())
	}

	return nil
}

func addEnvironmentVariables(currentPath string) error {
	filePath := fmt.Sprintf("%s%sconfig%senvironment.go", currentPath, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	renderer := domain.GetRenderer()
	err := renderer.BackupExistingCode(filePath)
	if err != nil {
		fmt.Printf("An error occurred while trying to backup the environment file. Error: %s\n", err.Error())
		return err
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("An error occurred while trying to read the environment file. Error: %s\n", err.Error())
		return err
	}

	code := strings.ReplaceAll(
		string(content),
		"type Config struct {\n",
		"type Config struct {\n//RedisAddress holds the address of the Redis cluster\nRedisAddress string\n//RedisPassword holds the password for the Redis cluster\nRedisPassword string\n")
	code = strings.ReplaceAll(
		code,
		"func init() {\n",
		"func init() {\n_ = viper.BindEnv(\"RedisAddress\", \"REDIS_ADDRESS\")\n_ = viper.BindEnv(\"RedisPassword\", \"REDIS_PASSWORD\")\n")

	err = ioutil.WriteFile(filePath, []byte(code), 0644)
	if err != nil {
		fmt.Printf("An error occurred while trying to update the environment file. Error: %s\n", err.Error())
		return err
	}

	return nil
}

//RenderredisinitTemplate renders the templates defined to the redisinit command with the proper variables
func RenderredisinitTemplate(args []string) error {
	spell := appcontext.Current.Get(appcontext.Spell).(tooldomain.Spell)
	renderer := domain.GetRenderer()
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("An error occurred while trying to init redis. Error: %s\n", err.Error())
		return err
	}
	moduleName := toolconfig.GetModuleName(currentPath)
	globalVariables := map[string]interface{}{
		"ModuleName": moduleName,
	}

	err = renderer.RenderTemplate(spell, "redisinit", globalVariables, nil)
	if err != nil {
		fmt.Printf("An error occurred while trying to render the template. Error: %s\n", err.Error())
		return err
	}

	err = addComponentConstantToContext(
		currentPath,
		"Cache = \"Cache\"")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = addComponentConstantToContext(
		currentPath,
		"Lock = \"Lock\"")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = addImportToMain(
		currentPath,
		fmt.Sprintf("%s/gateway/redis", moduleName))
	if err != nil {
		fmt.Printf("An error occurred while trying to add the import to main. Error: %s\n", err.Error())
		return err
	}

	err = addEnvironmentVariables(currentPath)
	if err != nil {
		fmt.Printf("An error occurred while trying to configure the environment. Error: %s\n", err.Error())
		return err
	}

	return nil
}
