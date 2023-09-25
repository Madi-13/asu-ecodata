package middleware

import (
	"app/domain"
	"app/initializer"
	"bytes"
	"context"
	"errors"
	"io/ioutil"

	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	PermissionScope = "Permission-Scope"
	Authorization   = "Authorization"
	ImportUserQuery = `select * from mdm_core.func_import_users($1)`
)

func CheckToken(c *fiber.Ctx) error {

	config, _ := initializer.LoadConfig(".")
	url := config.AuthServiceUrl

	authClient := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", url, "auth/validate"), nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	req.Header.Set(Authorization, c.Get(Authorization))

	res, err := authClient.Do(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err})
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err})
	}
	res.Body.Close()

	if res.StatusCode != fiber.StatusOK /*&& !strings.Contains(string(b), "token valid")*/ {
		c.Status(res.StatusCode).Set("Content-Type", "application/json")
		return c.Send(b)
	}

	auth := domain.Auth{}
	err = json.Unmarshal(b, &auth)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err})
	}
	c.Locals("X-User-Id", auth.Sub)
	//TODO List of roles
	c.Locals("Role", auth.RealmAccess.Roles[0])

	return c.Next()
}

func CheckTokenLC(authorization string) (int, string, string, error) {
	config, _ := initializer.LoadConfig(".")
	url := config.AuthServiceUrl

	authClient := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", url, "auth/validate"), nil)
	if err != nil {
		return 0, "", "", err
	}

	req.Header.Set("Authorization", authorization)

	res, err := authClient.Do(req)
	if err != nil {
		return res.StatusCode, "", "", err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, "", "", err
	}
	res.Body.Close()

	if res.StatusCode != fiber.StatusOK /*&& !strings.Contains(string(b), "token valid")*/ {
		return res.StatusCode, "", "", errors.New(string(b))
	}

	auth := domain.Auth{}
	err = json.Unmarshal(b, &auth)
	if err != nil {
		return res.StatusCode, "", "", err
	}

	return res.StatusCode, auth.Sub, auth.RealmAccess.Roles[0], nil
}

func SetRead(c *fiber.Ctx) error {
	c.Locals(PermissionScope, "read")
	return c.Next()
}

func SetWrite(c *fiber.Ctx) error {
	c.Locals(PermissionScope, "write")
	return c.Next()
}

func SetDelete(c *fiber.Ctx) error {
	c.Locals(PermissionScope, "delete")
	return c.Next()
}

func SetUpdate(c *fiber.Ctx) error {
	c.Locals(PermissionScope, "update")
	return c.Next()
}

func CheckPermission(c *fiber.Ctx) error {

	config, _ := initializer.LoadConfig(".")
	url := config.AuthServiceUrl

	authClient := &http.Client{Timeout: 10 * time.Second}

	payload := domain.GetRecordsInput{}
	err := json.Unmarshal(c.Body(), &payload)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err})
	}

	permission := domain.Permission{
		PermissionCode: payload.TableName,
		Scope:          fmt.Sprint(c.Locals(PermissionScope)),
	}

	body, err := json.Marshal(permission)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err})
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s%s", url, "auth/permission/check"),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	req.Header.Set(Authorization, c.Get(Authorization))

	res, err := authClient.Do(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err})
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err})
	}
	res.Body.Close()

	permissionRes := map[string]bool{}
	json.Unmarshal(b, &permissionRes)
	if res.StatusCode == fiber.StatusOK && permissionRes["has_permission"] == false {
		c.Status(http.StatusForbidden).Set("Content-Type", "application/json")
		return c.JSON(fiber.Map{"db_code": 403, "db_error_desc": "Access Denied", "result": make([]string, 0)})
	}

	return c.Next()
}

func ImportUsers() error {

	log.Println("ImportUsers")
	config, _ := initializer.LoadConfig(".")
	url := config.AuthServiceUrl

	authClient := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", url, "auth/login"), nil)
	if err != nil {
		return err
	}

	bscr := "Basic " + b64.StdEncoding.EncodeToString([]byte(config.Keycloak.User+":"+config.Keycloak.Password))

	req.Header.Set(Authorization, bscr)

	res, err := authClient.Do(req)
	if err != nil {
		return err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	res.Body.Close()

	auth := domain.TokenStruct{}
	err = json.Unmarshal(b, &auth)
	if err != nil {
		return err
	}

	reqa, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", url, "user/get-list"), nil)
	if err != nil {
		return err
	}

	reqa.Header.Set(Authorization, auth.TokenType+" "+auth.AccessToken)

	resa, err := authClient.Do(reqa)
	if err != nil {
		return err
	}

	ba, err := io.ReadAll(resa.Body)
	if err != nil {
		return err
	}
	res.Body.Close()

	auth_users := make([]domain.Users, 0)
	err = json.Unmarshal(ba, &auth_users)
	if err != nil {
		return err
	}

	con, err := initializer.GetConnection()
	if err != nil {
		return err
	}
	defer con.Release()
	out := domain.GetRecordsOut{}

	for _, auser := range auth_users {

		user := domain.DMLUsers{OperationType: "add", Users: domain.Users{ID: auser.ID, Username: auser.Username, FirstName: auser.FirstName, LastName: auser.LastName}}

		json_user, err := json.Marshal(user)
		if err != nil {
			return err
		}

		err = con.QueryRow(context.Background(), ImportUserQuery, json_user).Scan(&out.DbCode, &out.DbErrorDesc)
		if err != nil {
			return err
		}
	}

	return nil

}

func HttpError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(message))

}
func HttpGood(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(message))
}

func ReadBody(r http.Request, payload interface{}) interface{} {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(reqBody, &payload)
	if err != nil {
		return err
	}
	return nil

}
