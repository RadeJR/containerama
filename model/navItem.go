package model

type NavItem struct {
	Name  string
	Icon  string
	Route string
	Roles []string
}

var NavItems = [7]NavItem{
	{
		Name:  "Containers",
		Icon:  "fa-solid fa-box",
		Route: "/containers",
	},
	{
		Name:  "Networks",
		Icon:  "fa-solid fa-network-wired",
		Route: "/networks",
	},
	{
		Name:  "Secrets",
		Icon:  "fa-solid fa-key",
		Route: "/secrets",
	},
	{
		Name:  "Stacks",
		Icon:  "fa-solid fa-layer-group",
		Route: "/stacks",
	},
	{
		Name:  "Users",
		Icon:  "fa-solid fa-users",
		Route: "/users",
		Roles: []string{"admin"},
	},
	{
		Name:  "Settings",
		Icon:  "fa-solid fa-sliders",
		Route: "/settings",
		Roles: []string{"admin"},
	},
	{
		Name:  "Account",
		Icon:  "fa-solid fa-user",
		Route: "/account",
	},
}
