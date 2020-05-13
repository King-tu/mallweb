package utils

import (
	"strings"

	"encoding/json"

	"github.com/micro/go-micro/v2/metadata"
	"golang.org/x/net/context"
)

func HasPerm(ctx context.Context, perm string) bool {
	md, _ := metadata.FromContext(ctx)
	res, ok := md["Perms"]
	if !ok {
		return false
	}

	perms := strings.Split(res, ",")
	if Contains(perm, strings.Split(res, ",")) {
		return true
	}

	for _, _perm := range perms {
		if strings.Index(_perm, perm) == 0 {
			return true
		}
	}

	return false
}

func HasAction(ctx context.Context, perm, action string) bool {
	if !HasPerm(ctx, perm) {
		return false
	}

	md, _ := metadata.FromContext(ctx)
	res, ok := md["Perm-Actions"]
	if !ok {
		return false
	}

	actionsMap := make(map[string]string)
	json.Unmarshal([]byte(res), &actionsMap)

	actions, ok := actionsMap[perm]
	if !ok {
		return false
	}

	_actions := strings.Split(actions, ",")
	return Contains(action, _actions)
}

func HasRole(ctx context.Context, target string) bool {
	md, _ := metadata.FromContext(ctx)
	res, ok := md["Roles"]
	if !ok {
		return false
	}

	roles := strings.Split(res, ",")
	if Contains(target, strings.Split(res, ",")) {
		return true
	}

	for _, _role := range roles {
		if strings.Index(_role, target) == 0 {
			return true
		}
	}

	return false
}

func HasPermV2(md map[string]string, perm string) bool {
	res, ok := md["Perms"]
	if !ok {
		return false
	}

	perms := strings.Split(res, ",")
	if Contains(perm, strings.Split(res, ",")) {
		return true
	}

	for _, _perm := range perms {
		if strings.Index(_perm, perm) == 0 {
			return true
		}
	}

	return false
}

func HasRoleV2(md map[string]string, target string) bool {
	res, ok := md["Roles"]
	if !ok {
		return false
	}

	roles := strings.Split(res, ",")
	if Contains(target, strings.Split(res, ",")) {
		return true
	}

	for _, _role := range roles {
		if strings.Index(_role, target) == 0 {
			return true
		}
	}

	return false
}

func HasSuperAdminPerm(ctx context.Context, target string) bool {
	md, _ := metadata.FromContext(ctx)
	res, ok := md["Scope"]
	if !ok {
		return false
	}

	scopes := strings.Split(res, ",")
	for _, scope := range scopes {
		kv := strings.Split(scope, ":")
		if len(kv) != 2 {
			continue
		}

		if kv[0] != "super_admin" {
			continue
		}

		if kv[1] == target || kv[1] == "all" {
			return true
		}
	}

	return false
}
